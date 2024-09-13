package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"codnect.io/chrono"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"github.com/joho/godotenv"
	pb "github.com/misterdelle/miner-and-commander/pb/github.com/braiins/bos-plus-api/braiins/bos"
	pbV1 "github.com/misterdelle/miner-and-commander/pb/github.com/braiins/bos-plus-api/braiins/bos/v1"
)

// Embed the entire directory.
//
//go:embed templates
var templates embed.FS

// var ErrGetMinerStats = errors.New("error getting miner stats")
// var BOSminerNotRunningError = errors.New("BOSminer is not running")
// var BOSminerAPIConnectionError = errors.New("BOSminer API connection error: Connection refused (os error 111)")
// var BOSminerAPIConnectionError = errors.New("os error 111")

type ErrGetMinerStats struct {
	s string
}

type AuthenticationToken struct {
	Key      string
	Value    string
	TimeOutS int
}

type Config struct {
	Env     string
	WebPort string

	MinerAddress  string
	MinerUsername string
	MinerPassword string

	StartTimer bool
	fasceMap   map[string]uint64

	EMailSend         bool
	EMailSMTPHost     string
	EMailSMTPPort     int
	EMailSMTPUsername string
	EMailSMTPPassword string

	EMailFrom    string
	EMailTo      string
	EMailSubject string
	Mailer       Mail

	AuthToken AuthenticationToken
}

var app Config

var ctx context.Context
var authCtx context.Context

var authClient pbV1.AuthenticationServiceClient
var configClient pbV1.ConfigurationServiceClient
var performanceClient pbV1.PerformanceServiceClient
var actionsClient pbV1.ActionsServiceClient
var apiVersionClient pb.ApiVersionServiceClient
var minerServiceClient pbV1.MinerServiceClient

func init() {
	//
	// Set up config
	//
	app = Config{}

	profilePassed := false
	flag.StringVar(&app.Env, "env", "", "Environment (SVIL|PROD")

	if app.Env != "" {
		profilePassed = true
		godotenv.Load(".env." + app.Env + ".local")
		godotenv.Load(".env." + app.Env)
	} else {
		godotenv.Load() // The Original .env
		app.Env = os.Getenv("Env")
	}

	if profilePassed {
		log.Printf("app.Env: %s", app.Env)
	} else {
		log.Println("app.Env NON settato, carico i dati dal file .env")
		log.Printf("app.Env: %s", app.Env)
	}

	app.WebPort = os.Getenv("WebPort")

	app.fasceMap = make(map[string]uint64, 5)

	//
	// Configurazione Miner
	//
	app.MinerAddress = os.Getenv("MinerAddress")
	app.MinerUsername = os.Getenv("MinerUsername")
	app.MinerPassword = os.Getenv("MinerPassword")

	app.StartTimer, _ = strconv.ParseBool(os.Getenv("StartTimer"))

	// Thresholds configuration
	fascia1PowerThreshold, _ := strconv.Atoi(os.Getenv("Fascia1PowerThreshold"))
	fascia2PowerThreshold, _ := strconv.Atoi(os.Getenv("Fascia2PowerThreshold"))
	fascia3PowerThreshold, _ := strconv.Atoi(os.Getenv("Fascia3PowerThreshold"))
	fascia4PowerThreshold, _ := strconv.Atoi(os.Getenv("Fascia4PowerThreshold"))
	fascia5PowerThreshold, _ := strconv.Atoi(os.Getenv("Fascia5PowerThreshold"))

	app.fasceMap[os.Getenv("Fascia1StartCronTime")] = uint64(fascia1PowerThreshold)
	app.fasceMap[os.Getenv("Fascia2StartCronTime")] = uint64(fascia2PowerThreshold)
	app.fasceMap[os.Getenv("Fascia3StartCronTime")] = uint64(fascia3PowerThreshold)
	app.fasceMap[os.Getenv("Fascia4StartCronTime")] = uint64(fascia4PowerThreshold)
	app.fasceMap[os.Getenv("Fascia5StartCronTime")] = uint64(fascia5PowerThreshold)

	//
	// Configurazione Mailer
	//
	app.EMailSend, _ = strconv.ParseBool(os.Getenv("EMailSend"))
	app.EMailSMTPHost = os.Getenv("EMailSMTPHost")
	app.EMailSMTPPort, _ = strconv.Atoi(os.Getenv("EMailSMTPPort"))
	app.EMailSMTPUsername = os.Getenv("EMailSMTPUsername")
	app.EMailSMTPPassword = os.Getenv("EMailSMTPPassword")
	app.EMailFrom = os.Getenv("EMailFrom")
	app.EMailTo = os.Getenv("EMailTo")
	app.EMailSubject = os.Getenv("EMailSubject")

	app.Mailer = app.createMailer()

	//
	// Set up Authentication Token
	//
	authToken := AuthenticationToken{
		Key:      "authorization",
		Value:    "",
		TimeOutS: 0,
	}
	app.AuthToken = authToken

}

func main() {

	retryOpts := []retry.CallOption{
		retry.WithMax(1_000),
		retry.WithBackoff(retry.BackoffExponential(100 * time.Millisecond)),
		retry.WithCodes(codes.Unavailable),
		retry.WithOnRetryCallback(OnRetryCallback),
	}

	credsOpt := grpc.WithTransportCredentials(insecure.NewCredentials())

	opts := []grpc.DialOption{
		credsOpt,
		grpc.WithChainUnaryInterceptor(
			retry.UnaryClientInterceptor(retryOpts...),
			unaryAuthInterceptor,
		),
	}

	// Set up a connection to the server.
	conn, err := grpc.NewClient(app.MinerAddress, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	authClient = pbV1.NewAuthenticationServiceClient(conn)
	configClient = pbV1.NewConfigurationServiceClient(conn)
	performanceClient = pbV1.NewPerformanceServiceClient(conn)
	actionsClient = pbV1.NewActionsServiceClient(conn)
	apiVersionClient = pb.NewApiVersionServiceClient(conn)
	minerServiceClient = pbV1.NewMinerServiceClient(conn)

	ctx = context.Background()

	// Get Miner Firmware Version
	// Non è necessaria la login
	log.Println("Fetching miner firmware version")
	apiVersion, err := GetAPIVersion(ctx)
	if err != nil {
		log.Fatalf("could not get miner firmware version: %v", err)
	}
	// log.Println("apiVersion: ", apiVersion.(*pb.ApiVersion))
	log.Println("apiVersion: ", apiVersion)

	err = app.LoginWrapper()
	if err != nil {
		log.Fatalf("could not login: %v", err)
	}

	// Listen for signals
	go app.listenForShutdown()

	// Schedules the check and eventually refresh of the authentication token
	taskScheduler := chrono.NewDefaultTaskScheduler()

	secondsToWait := time.Duration(app.AuthToken.TimeOutS) * time.Second

	_, err = taskScheduler.ScheduleWithFixedDelay(func(ctx context.Context) {
		log.Printf("Fixed Delay Task every %s", secondsToWait)

		if app.AuthToken.Value != "" {
			log.Print("Starting refresh authentication token...")

			err = app.LoginWrapper()
			if err != nil {
				log.Fatalf("could not login: %v", err)
			}

			log.Print("Ending refresh authentication token...")
		}
	}, secondsToWait)
	if err == nil {
		log.Print("Task has been scheduled successfully.")
	}
	// log.Print("task: ", task)

	if app.StartTimer {
		// taskScheduler := chrono.NewDefaultTaskScheduler()

		for k, v := range app.fasceMap {
			startCronTime := k
			powerThreshold := v

			if startCronTime != "-" {
				taskFascia, err := taskScheduler.ScheduleWithCron(func(ctx context.Context) {
					log.Printf("Scheduled Task With Cron: %v", powerThreshold)

					var msgBody []string

					if powerThreshold == 0 {
						/*
						* Se la powerThreshold è uguale zero spengo il miner
						 */
						log.Println("Stopping miner")

						_, err := MinerStop(authCtx)
						if err != nil {
							log.Println("could not stop miner", err)
						}

						msgBody = append(msgBody, "Stopped miner")
					} else {
						/*
						* Se la powerThreshold è maggiore di zero nel dubbio lo faccio partire e poi
						* setto la powerThreshold sul miner
						 */

						log.Println("Starting miner")

						_, err := MinerStart(authCtx)
						if err != nil {
							log.Println("could not start miner", err)
						}

						log.Println("Setting Power Target to ", powerThreshold)

						_, err = MinerSetPowerTarget(authCtx, powerThreshold)
						if err != nil {
							log.Printf("could not set power target: %v", err)
						}

						msgBody = append(msgBody, fmt.Sprintf("Set Power Target to %v", powerThreshold))
					}

					app.sendEMail(msgBody)
				}, startCronTime)
				if err != nil {
					log.Printf("Errore: %s", err)
					return
				}

				taskFascia.IsCancelled()
			}
		}

	}

	//
	// listen for web connections
	//
	app.serve()

}

func (app *Config) listenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	app.shutdown()
	os.Exit(0)
}

func (app *Config) shutdown() {
	// perform any cleanup tasks
	log.Println("Shutting down application...")
}

func (app *Config) serve() {
	//
	// start http server
	//
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.WebPort),
		Handler: app.routes(),
	}

	log.Printf("The WEB server is listening on: 0.0.0.0:%s\n", app.WebPort)
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func OnRetryCallback(ctx context.Context, attempt uint, err error) {
	log.Printf("grpc_retry attempt: %d, backoff for %v", attempt, err)
}

func (e ErrGetMinerStats) Error() string {
	return e.s
}

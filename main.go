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
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	"github.com/joho/godotenv"
	pb "github.com/misterdelle/miner-and-commander/pb/github.com/braiins/bos-plus-api/braiins/bos"
	pbV1 "github.com/misterdelle/miner-and-commander/pb/github.com/braiins/bos-plus-api/braiins/bos/v1"
)

// Embed the entire directory.
//
//go:embed templates
var templates embed.FS

type Config struct {
	Env     string
	WebPort string

	MinerAddress  string
	MinerUsername string
	MinerPassword string

	fasceMap map[string]uint64

	EMailSend         bool
	EMailSMTPHost     string
	EMailSMTPPort     int
	EMailSMTPUsername string
	EMailSMTPPassword string

	EMailFrom    string
	EMailTo      string
	EMailSubject string
	Mailer       Mail
}

var app Config

var ctx context.Context
var authCtx context.Context

// var headerMD metadata.MD

var authClient pbV1.AuthenticationServiceClient
var configClient pbV1.ConfigurationServiceClient
var performanceClient pbV1.PerformanceServiceClient
var actionsClient pbV1.ActionsServiceClient
var apiVersionClient pb.ApiVersionServiceClient

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

}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.NewClient(app.MinerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	authClient = pbV1.NewAuthenticationServiceClient(conn)
	configClient = pbV1.NewConfigurationServiceClient(conn)
	performanceClient = pbV1.NewPerformanceServiceClient(conn)
	actionsClient = pbV1.NewActionsServiceClient(conn)
	apiVersionClient = pb.NewApiVersionServiceClient(conn)

	ctx = context.Background()
	headerMD := metadata.MD{}

	_, err = RetryWithBackoff(Login, 1_000_000, 2*time.Second, &headerMD)
	if err != nil {
		log.Fatalf("could not login: %v", err)
	} else {
		log.Println("Login successful")
	}

	authTokens := headerMD.Get("authorization")
	if len(authTokens) == 0 {
		log.Fatal("authorization token not found in headers")
	}
	authToken := authTokens[0] // Taking the first token

	log.Println("authToken: ", authToken)

	// Attach auth token to context
	md := metadata.New(map[string]string{"authorization": authToken})
	authCtx = metadata.NewOutgoingContext(ctx, md)

	// Get Miner Firmware Version
	log.Println("Fetching miner firmware version")
	apiVersion, err := RetryWithBackoff(GetAPIVersion, 1_000_000, 2*time.Second, authCtx)
	if err != nil {
		log.Fatalf("could not get miner firmware version: %v", err)
	}
	log.Println("apiVersion: ", apiVersion.(*pb.ApiVersion))

	// Get Miner Configuration
	log.Println("Fetching miner configuration")
	minerConfigResponse, err := RetryWithBackoff(GetMinerConfiguration, 1_000_000, 2*time.Second, authCtx)
	if err != nil {
		log.Fatalf("could not get miner configuration: %v", err)
	}
	log.Println("minerConfigResponse: ", minerConfigResponse)

	// Listen for signals
	go app.listenForShutdown()

	taskScheduler := chrono.NewDefaultTaskScheduler()

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

					_, err := RetryWithBackoff(MinerStop, 1_000_000, 2*time.Second, authCtx)
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

					_, err := RetryWithBackoff(MinerStart, 1_000_000, 2*time.Second, authCtx)
					if err != nil {
						log.Println("could not start miner", err)
					}

					log.Println("Setting Power Target to ", powerThreshold)

					_, err = RetryWithBackoff(MinerSetPowerTarget, 1_000_000, 2*time.Second, authCtx, powerThreshold)
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

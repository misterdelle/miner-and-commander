package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"codnect.io/chrono"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	"github.com/joho/godotenv"
	pb "github.com/misterdelle/miner-and-commander/pb/github.com/braiins/bos-plus-api/braiins/bos/v1"
)

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

	LogFile string
}

var app Config

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

	app.LogFile = os.Getenv("LogFile")

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
	// conn, err := grpc.Dial(app.MinerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	conn, err := grpc.NewClient(app.MinerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	authClient := pb.NewAuthenticationServiceClient(conn)
	configClient := pb.NewConfigurationServiceClient(conn)
	performanceClient := pb.NewPerformanceServiceClient(conn)
	actionsClient := pb.NewActionsServiceClient(conn)

	// Login
	log.Println("Logging in...")
	headerMD := metadata.MD{}
	ctx := context.Background()
	loginReq := &pb.LoginRequest{
		Username: app.MinerUsername,
		Password: app.MinerPassword,
	}

	_, err = authClient.Login(ctx, loginReq, grpc.Header(&headerMD))
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
	authCtx := metadata.NewOutgoingContext(ctx, md)

	// Get Miner Configuration
	log.Println("Fetching miner configuration")
	minerConfigResponse, err := configClient.GetMinerConfiguration(authCtx, &pb.GetMinerConfigurationRequest{})
	if err != nil {
		log.Fatalf("could not get miner configuration: %v", err)
	}

	log.Println("minerConfigResponse: ", minerConfigResponse)

	// fmt.Println("Trying to pause mining for 10 seconds...")
	// // Pause and Resume Mining
	// _, err = actionsClient.PauseMining(authCtx, &pb.PauseMiningRequest{})
	// if err != nil {
	// 	log.Fatalf("could not pause mining: %v", err)
	// }

	// time.Sleep(10 * time.Second)

	// fmt.Println("Resuming mining")
	// _, err = actionsClient.ResumeMining(authCtx, &pb.ResumeMiningRequest{})
	// if err != nil {
	// 	log.Fatalf("could not resume mining: %v", err)
	// }

	// time.Sleep(5 * time.Second)

	// log.Println("Operation successful")

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

					_, err = actionsClient.Stop(authCtx, &pb.StopRequest{})
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

					_, err = actionsClient.Start(authCtx, &pb.StartRequest{})
					if err != nil {
						log.Println("could not start miner", err)
					}

					log.Println("Setting Power Target to ", powerThreshold)

					_, err = performanceClient.SetPowerTarget(authCtx, &pb.SetPowerTargetRequest{
						SaveAction: pb.SaveAction_SAVE_ACTION_SAVE_AND_APPLY,
						PowerTarget: &pb.Power{
							Watt: powerThreshold,
						},
					})
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

	// if app.Fascia1StartCronTime != "-" {
	// 	taskFascia1, err := taskScheduler.ScheduleWithCron(func(ctx context.Context) {
	// 		logger.Logger.Info(fmt.Sprintf("Scheduled Task With Cron: %s", app.Fascia1StartCronTime))

	// 		log.Println("Setting Power Target to ", app.Fascia1PowerThreshold)
	// 		_, err = performanceClient.SetPowerTarget(authCtx, &pb.SetPowerTargetRequest{
	// 			SaveAction: pb.SaveAction_SAVE_ACTION_SAVE_AND_APPLY,
	// 			PowerTarget: &pb.Power{
	// 				Watt: uint64(app.Fascia1PowerThreshold),
	// 			},
	// 		})
	// 		if err != nil {
	// 			log.Fatalf("could not set power target: %v", err)
	// 		}

	// 	}, app.Fascia1StartCronTime)
	// 	if err != nil {
	// 		logger.Logger.Error(fmt.Sprintf("Errore: %s", err))
	// 		return
	// 	}

	// 	taskFascia1.IsCancelled()
	// }

	// if app.Fascia2StartCronTime != "-" {
	// 	taskFascia2, err := taskScheduler.ScheduleWithCron(func(ctx context.Context) {
	// 		logger.Logger.Info(fmt.Sprintf("Scheduled Task With Cron: %s", app.Fascia2StartCronTime))

	// 		log.Println("Setting Power Target to ", app.Fascia2PowerThreshold)
	// 		_, err = performanceClient.SetPowerTarget(authCtx, &pb.SetPowerTargetRequest{
	// 			SaveAction: pb.SaveAction_SAVE_ACTION_SAVE_AND_APPLY,
	// 			PowerTarget: &pb.Power{
	// 				Watt: uint64(app.Fascia2PowerThreshold),
	// 			},
	// 		})
	// 		if err != nil {
	// 			log.Fatalf("could not set power target: %v", err)
	// 		}

	// 	}, app.Fascia2StartCronTime)
	// 	if err != nil {
	// 		logger.Logger.Error(fmt.Sprintf("Errore: %s", err))
	// 		return
	// 	}

	// 	taskFascia2.IsCancelled()
	// }

	// if app.Fascia3StartCronTime != "-" {
	// 	taskFascia3, err := taskScheduler.ScheduleWithCron(func(ctx context.Context) {
	// 		logger.Logger.Info(fmt.Sprintf("Scheduled Task With Cron: %s", app.Fascia3StartCronTime))

	// 		log.Println("Setting Power Target to ", app.Fascia3PowerThreshold)
	// 		_, err = performanceClient.SetPowerTarget(authCtx, &pb.SetPowerTargetRequest{
	// 			SaveAction: pb.SaveAction_SAVE_ACTION_SAVE_AND_APPLY,
	// 			PowerTarget: &pb.Power{
	// 				Watt: uint64(app.Fascia3PowerThreshold),
	// 			},
	// 		})
	// 		if err != nil {
	// 			log.Fatalf("could not set power target: %v", err)
	// 		}

	// 	}, app.Fascia3StartCronTime)
	// 	if err != nil {
	// 		logger.Logger.Error(fmt.Sprintf("Errore: %s", err))
	// 		return
	// 	}

	// 	taskFascia3.IsCancelled()
	// }

	// if app.Fascia4StartCronTime != "-" {
	// 	taskFascia4, err := taskScheduler.ScheduleWithCron(func(ctx context.Context) {
	// 		logger.Logger.Info(fmt.Sprintf("Scheduled Task With Cron: %s", app.Fascia4StartCronTime))

	// 		log.Println("Setting Power Target to ", app.Fascia4PowerThreshold)
	// 		_, err = performanceClient.SetPowerTarget(authCtx, &pb.SetPowerTargetRequest{
	// 			SaveAction: pb.SaveAction_SAVE_ACTION_SAVE_AND_APPLY,
	// 			PowerTarget: &pb.Power{
	// 				Watt: uint64(app.Fascia4PowerThreshold),
	// 			},
	// 		})
	// 		if err != nil {
	// 			log.Fatalf("could not set power target: %v", err)
	// 		}

	// 	}, app.Fascia4StartCronTime)
	// 	if err != nil {
	// 		logger.Logger.Error(fmt.Sprintf("Errore: %s", err))
	// 		return
	// 	}

	// 	taskFascia4.IsCancelled()
	// }

	// if app.Fascia5StartCronTime != "-" {
	// 	taskFascia5, err := taskScheduler.ScheduleWithCron(func(ctx context.Context) {
	// 		logger.Logger.Info(fmt.Sprintf("Scheduled Task With Cron: %s", app.Fascia5StartCronTime))

	// 		log.Println("Setting Power Target to ", app.Fascia5PowerThreshold)
	// 		_, err = performanceClient.SetPowerTarget(authCtx, &pb.SetPowerTargetRequest{
	// 			SaveAction: pb.SaveAction_SAVE_ACTION_SAVE_AND_APPLY,
	// 			PowerTarget: &pb.Power{
	// 				Watt: uint64(app.Fascia5PowerThreshold),
	// 			},
	// 		})
	// 		if err != nil {
	// 			log.Fatalf("could not set power target: %v", err)
	// 		}

	// 	}, app.Fascia5StartCronTime)
	// 	if err != nil {
	// 		logger.Logger.Error(fmt.Sprintf("Errore: %s", err))
	// 		return
	// 	}

	// 	taskFascia5.IsCancelled()
	// }

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

package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	cb "github.com/emirpasic/gods/queues/circularbuffer"
	"github.com/misterdelle/miner-and-commander/pkg/miner_ops"
	"github.com/misterdelle/miner-and-commander/pkg/model"
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

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Embed the entire directory.
//
//go:embed templates
var templates embed.FS

var stationData = model.NewStation()

var minerConfigurations = map[string]model.MinerConfiguration{}
var pvRead = cb.New(30)

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

	StartTimer             bool
	TimerIntervalinMinutes int
	//fasceMap               map[string]uint64

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

	MQTTURL       string
	MQTTUser      string
	MQTTPassword  string
	MQTTTopicName string
	MQTTClient    mqtt.Client

	MinerOperations *miner_ops.MinerOps

	MinerThresholdList []*model.MinerThreshold
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

	//app.fasceMap = make(map[string]uint64, 5)

	//
	// Configurazione Miner
	//
	app.MinerAddress = os.Getenv("MinerAddress")
	app.MinerUsername = os.Getenv("MinerUsername")
	app.MinerPassword = os.Getenv("MinerPassword")

	app.StartTimer, _ = strconv.ParseBool(os.Getenv("StartTimer"))
	app.TimerIntervalinMinutes, _ = strconv.Atoi(os.Getenv("TimerIntervalinMinutes"))

	// Thresholds configuration
	//fascia1PowerThreshold, _ := strconv.Atoi(os.Getenv("Fascia1PowerThreshold"))
	//fascia2PowerThreshold, _ := strconv.Atoi(os.Getenv("Fascia2PowerThreshold"))
	//fascia3PowerThreshold, _ := strconv.Atoi(os.Getenv("Fascia3PowerThreshold"))
	//fascia4PowerThreshold, _ := strconv.Atoi(os.Getenv("Fascia4PowerThreshold"))
	//fascia5PowerThreshold, _ := strconv.Atoi(os.Getenv("Fascia5PowerThreshold"))

	//app.fasceMap[os.Getenv("Fascia1StartCronTime")] = uint64(fascia1PowerThreshold)
	//app.fasceMap[os.Getenv("Fascia2StartCronTime")] = uint64(fascia2PowerThreshold)
	//app.fasceMap[os.Getenv("Fascia3StartCronTime")] = uint64(fascia3PowerThreshold)
	//app.fasceMap[os.Getenv("Fascia4StartCronTime")] = uint64(fascia4PowerThreshold)
	//app.fasceMap[os.Getenv("Fascia5StartCronTime")] = uint64(fascia5PowerThreshold)

	var mt1 *model.MinerThreshold = model.ParseMinerThreshold(os.Getenv("MinerThreshold1"))
	var mt2 *model.MinerThreshold = model.ParseMinerThreshold(os.Getenv("MinerThreshold2"))
	var mt3 *model.MinerThreshold = model.ParseMinerThreshold(os.Getenv("MinerThreshold3"))
	var mt4 *model.MinerThreshold = model.ParseMinerThreshold(os.Getenv("MinerThreshold4"))
	var mt5 *model.MinerThreshold = model.ParseMinerThreshold(os.Getenv("MinerThreshold5"))
	var mt6 *model.MinerThreshold = model.ParseMinerThreshold(os.Getenv("MinerThreshold6"))
	var mt7 *model.MinerThreshold = model.ParseMinerThreshold(os.Getenv("MinerThreshold7"))
	var mt8 *model.MinerThreshold = model.ParseMinerThreshold(os.Getenv("MinerThreshold8"))
	var mt9 *model.MinerThreshold = model.ParseMinerThreshold(os.Getenv("MinerThreshold9"))

	app.MinerThresholdList = append(app.MinerThresholdList, mt1)
	app.MinerThresholdList = append(app.MinerThresholdList, mt2)
	app.MinerThresholdList = append(app.MinerThresholdList, mt3)
	app.MinerThresholdList = append(app.MinerThresholdList, mt4)
	app.MinerThresholdList = append(app.MinerThresholdList, mt5)
	app.MinerThresholdList = append(app.MinerThresholdList, mt6)
	app.MinerThresholdList = append(app.MinerThresholdList, mt7)
	app.MinerThresholdList = append(app.MinerThresholdList, mt8)
	app.MinerThresholdList = append(app.MinerThresholdList, mt9)

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

	//
	// Configurazione Server MQTT
	//
	app.MQTTURL = os.Getenv("MQTT.URL")
	app.MQTTUser = os.Getenv("MQTT.User")
	app.MQTTPassword = os.Getenv("MQTT.Password")
	app.MQTTTopicName = os.Getenv("MQTT.Prefix")

	//
	// Carico la mappa delle configurazioni del miner
	//
	model.LoadMinerConfigurationsMap(minerConfigurations)
}

func main() {

	var broker = app.MQTTURL
	mqttOpts := mqtt.NewClientOptions()
	mqttOpts.AddBroker(broker)
	mqttOpts.SetClientID("go_mqtt_client")
	mqttOpts.SetUsername(app.MQTTUser)
	mqttOpts.SetPassword(app.MQTTPassword)
	mqttOpts.SetConnectRetry(true)
	mqttOpts.OnConnect = connectHandler
	mqttOpts.OnConnectionLost = connectLostHandler
	app.MQTTClient = mqtt.NewClient(mqttOpts)

	if token := app.MQTTClient.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf(fmt.Sprintf("Error connecting to MQTT broker: %s", token.Error()))
	}

	go app.subscribeTopic()

	retryOpts := []retry.CallOption{
		retry.WithMax(10),
		retry.WithBackoff(retry.BackoffExponential(100 * time.Millisecond)),
		retry.WithCodes(codes.Unavailable),
		retry.WithOnRetryCallback(OnRetryCallback),
	}

	credsOpt := grpc.WithTransportCredentials(insecure.NewCredentials())

	grpcOpts := []grpc.DialOption{
		credsOpt,
		grpc.WithChainUnaryInterceptor(
			retry.UnaryClientInterceptor(retryOpts...),
			unaryAuthInterceptor,
		),
	}

	// Set up a connection to the server.
	conn, err := grpc.NewClient(app.MinerAddress, grpcOpts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx = context.Background()
	authCtx = ctx

	authClient = pbV1.NewAuthenticationServiceClient(conn)
	configClient = pbV1.NewConfigurationServiceClient(conn)
	performanceClient = pbV1.NewPerformanceServiceClient(conn)
	actionsClient = pbV1.NewActionsServiceClient(conn)
	apiVersionClient = pb.NewApiVersionServiceClient(conn)
	minerServiceClient = pbV1.NewMinerServiceClient(conn)

	app.MinerOperations = miner_ops.NewMinerOps(authCtx, apiVersionClient, configClient, actionsClient, minerServiceClient, performanceClient)

	// Get Miner Firmware Version
	// Non è necessaria la login
	log.Println("Fetching miner firmware version")
	apiVersion, err := app.MinerOperations.GetAPIVersion()
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

		/*
			for k, v := range app.fasceMap {
				startCronTime := k
				powerThreshold := v

				if startCronTime != "-" {
					taskFascia, err := taskScheduler.ScheduleWithCron(func(ctx context.Context) {
						log.Printf("Scheduled Task With Cron: %v", powerThreshold)

						var msgBody []string

						if powerThreshold == 0 {
							//
							// Se la powerThreshold è uguale zero spengo il miner
							//
							log.Println("Stopping miner")

							_, err := MinerStop(authCtx)
							if err != nil {
								log.Println("could not stop miner", err)
							}

							msgBody = append(msgBody, "Stopped miner")
						} else {
							//
							// Se la powerThreshold è maggiore di zero nel dubbio lo faccio partire e poi
							// setto la powerThreshold sul miner
							//

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
		*/

		ticker := time.NewTicker(time.Duration(app.TimerIntervalinMinutes) * time.Minute)

		go func() {
			for {
				select {
				case t := <-ticker.C:
					log.Printf("Tick at: %v\n", t)
					log.Println("app.startCheck()")
					app.startCheck()
				}
			}
		}()

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

	app.MQTTClient.Unsubscribe(app.MQTTTopicName)
	app.MQTTClient.Disconnect(250)

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

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Connected")
	go app.subscribeTopic()
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Printf("Connect lost: %v", err)
}

func (app *Config) subscribeTopic() {
	topicFilter := map[string]byte{
		app.MQTTTopicName + "/station/lastUpdateTime":          0,
		app.MQTTTopicName + "/station/totalPowerFromPV":        0,
		app.MQTTTopicName + "/station/currentConsumptionPower": 0,
		app.MQTTTopicName + "/station/batterySOC":              0,
		app.MQTTTopicName + "/station/batteryPower":            0,
		app.MQTTTopicName + "/LoadInfo/Load C/Load C Power":    0,
	}

	if token := app.MQTTClient.SubscribeMultiple(topicFilter, onMessageReceived); token.Wait() && token.Error() != nil {
		log.Fatalf(fmt.Sprintf("Error subscribing to topic: %s", token.Error()))
	}
}

func (app *Config) startCheck() {
	var batterySOCThreshold float32 = 80
	var msgBody []string

	//
	// Faccio partire il miner se:
	// 1) le batterie sono almeno al 80%
	// 2) Sto producendo tra i 500W ed i 1000W setto il miner a 300W con una sola Hashboard abilitata (settata random ad ogni check)
	// 3) Sto producendo tra i 1001W ed i 1500W setto il miner a 600W con due Hashboard abilitate (settate random ad ogni check)
	// 4) Sto producendo tra i 1501W ed i 2000W setto il miner a 1200W con tutte e tre le Hashboard abilitate
	// 5) Sto producendo tra i 2001W ed i 2500W setto il miner a 1500W con tutte e tre le Hashboard abilitate
	// 6) Sto producendo tra i 2501W ed i 3000W setto il miner a 2000W con tutte e tre le Hashboard abilitate
	// 7) Sto producendo tra i 3001W ed i 3500W setto il miner a 2500W con tutte e tre le Hashboard abilitate
	// 8) Sto producendo oltre i 3501w setto il miner a 3068W con tutte e tre le Hashboard abilitate
	//
	if stationData.CurrentBatterySOC >= batterySOCThreshold {
		totalPowerFromPV := stationData.CurrentTotalPowerFromPV

		minerConfigName := model.GetMinerConfigurationNameByThreshold(app.MinerThresholdList, uint64(totalPowerFromPV))
		minerConfig := minerConfigurations[minerConfigName]
		app.MinerOperations.SetMinerConfiguration(&minerConfig)
		msgBody = append(msgBody, fmt.Sprintf("Total Power from PV: %v Set Power Target to %v with Hashboards: %s", totalPowerFromPV, minerConfig.PowerThreshold, minerConfig.HashboardIds))

		//if totalPowerFromPV >= 500 && totalPowerFromPV <= 1000 {
		//	minerConfig := minerConfigurations["300"]
		//	app.MinerOperations.SetMinerConfiguration(&minerConfig)
		//	msgBody = append(msgBody, fmt.Sprintf("Set Power Target to %v with Hashboards: %s", minerConfig.PowerThreshold, minerConfig.HashboardIds))
		//} else if totalPowerFromPV >= 1001 && totalPowerFromPV <= 1500 {
		//	minerConfig := minerConfigurations["600"]
		//	app.MinerOperations.SetMinerConfiguration(&minerConfig)
		//	msgBody = append(msgBody, fmt.Sprintf("Set Power Target to %v with Hashboards: %s", minerConfig.PowerThreshold, minerConfig.HashboardIds))
		//} else if totalPowerFromPV >= 1501 && totalPowerFromPV <= 2000 {
		//	minerConfig := minerConfigurations["1200"]
		//	app.MinerOperations.SetMinerConfiguration(&minerConfig)
		//	msgBody = append(msgBody, fmt.Sprintf("Set Power Target to %v with Hashboards: %s", minerConfig.PowerThreshold, minerConfig.HashboardIds))
		//} else if totalPowerFromPV >= 2001 && totalPowerFromPV <= 2500 {
		//	minerConfig := minerConfigurations["1500"]
		//	app.MinerOperations.SetMinerConfiguration(&minerConfig)
		//	msgBody = append(msgBody, fmt.Sprintf("Set Power Target to %v with Hashboards: %s", minerConfig.PowerThreshold, minerConfig.HashboardIds))
		//} else if totalPowerFromPV >= 2501 && totalPowerFromPV <= 3000 {
		//	minerConfig := minerConfigurations["2000"]
		//	app.MinerOperations.SetMinerConfiguration(&minerConfig)
		//	msgBody = append(msgBody, fmt.Sprintf("Set Power Target to %v with Hashboards: %s", minerConfig.PowerThreshold, minerConfig.HashboardIds))
		//} else if totalPowerFromPV >= 3001 && totalPowerFromPV <= 3500 {
		//	minerConfig := minerConfigurations["2500"]
		//	app.MinerOperations.SetMinerConfiguration(&minerConfig)
		//	msgBody = append(msgBody, fmt.Sprintf("Set Power Target to %v with Hashboards: %s", minerConfig.PowerThreshold, minerConfig.HashboardIds))
		//} else if totalPowerFromPV >= 3501 {
		//	minerConfig := minerConfigurations["3068"]
		//	app.MinerOperations.SetMinerConfiguration(&minerConfig)
		//	msgBody = append(msgBody, fmt.Sprintf("Set Power Target to %v with Hashboards: %s", minerConfig.PowerThreshold, minerConfig.HashboardIds))
		//} else {
		//	msg := fmt.Sprintf("Batterie sopra al %.2f%%, ma produzione insufficiente, fermo il miner", batterySOCThreshold)
		//	log.Println(msg)
		//
		//	//
		//	// Nel dubbio fermo il miner
		//	//
		//	log.Println("Stopping miner")
		//
		//	_, err := app.MinerOperations.MinerStop()
		//	if err != nil {
		//		log.Println("could not stop miner", err)
		//	}
		//
		//	msgBody = append(msgBody, msg)
		//}
	} else {
		msg := fmt.Sprintf("Batterie sotto al %.2f%%, non faccio nulla", batterySOCThreshold)
		log.Println(msg)

		msgBody = append(msgBody, msg)
	}

	app.sendEMail(msgBody)
}

func (app *Config) applyConfig(minerConfigName string) {
	var msgBody []string

	minerConfig := minerConfigurations[minerConfigName]
	app.MinerOperations.SetMinerConfiguration(&minerConfig)
	msgBody = append(msgBody, fmt.Sprintf("Set Power Target to %v with Hashboards: %s", minerConfig.PowerThreshold, minerConfig.HashboardIds))

	app.sendEMail(msgBody)
}

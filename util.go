package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	pb "github.com/misterdelle/miner-and-commander/pb/github.com/braiins/bos-plus-api/braiins/bos"
	pbV1 "github.com/misterdelle/miner-and-commander/pb/github.com/braiins/bos-plus-api/braiins/bos/v1"
)

func (app *Config) createMailer() Mail {
	m := Mail{
		Domain:      "miner-and-commander.com",
		Host:        app.EMailSMTPHost,
		Port:        app.EMailSMTPPort,
		Username:    app.EMailSMTPUsername,
		Password:    app.EMailSMTPPassword,
		Encryption:  "tls",
		FromName:    "Miner And Commander",
		FromAddress: "MinerAndCommander@miner-and-commander.com",
	}

	return m
}

func (app *Config) sendEMail(msgBody []string) {
	if app.EMailSend {
		to := strings.Split(app.EMailTo, ";")

		jsonMsgBody, _ := json.MarshalIndent(msgBody, "", "\t")
		msgBodyOut := formatJSON([]byte(jsonMsgBody), &msgBody)

		msg := Message{
			From:       app.EMailFrom,
			To:         to,
			Subject:    app.EMailSubject,
			Data:       msgBodyOut,
			Encryption: "tls",
			//Attachments: "",
		}

		err := app.Mailer.SendSMTPMessage(msg)
		if err != nil {
			log.Printf("Errore: %s", err)
			return
		}

		log.Println("Mail sent!!!")
	}
}

func formatJSON(str []byte, any interface{}) string {
	err := json.Unmarshal(str, &any)
	if err != nil {
		panic(err)
	}

	formattedString, _ := json.MarshalIndent(&any, "", "  ")

	return string(formattedString)
}

// Get Miner Configuration
func GetMinerConfiguration(authCtx context.Context) (*pbV1.GetMinerConfigurationResponse, error) {
	minerConfigResponse, err := configClient.GetMinerConfiguration(authCtx, &pbV1.GetMinerConfigurationRequest{})
	if err != nil {
		return nil, err
	}

	return minerConfigResponse, nil
}

// Get Miner Details
func GetMinerDetails(authCtx context.Context) (*pbV1.GetMinerDetailsResponse, error) {
	minerDetailsResponse, err := minerServiceClient.GetMinerDetails(authCtx, &pbV1.GetMinerDetailsRequest{})
	if err != nil {
		return nil, err
	}

	return minerDetailsResponse, nil
}

// Get Miner Stats
func GetMinerStats(authCtx context.Context) (*pbV1.GetMinerStatsResponse, error) {
	minerStatsResponse, err := minerServiceClient.GetMinerStats(authCtx, &pbV1.GetMinerStatsRequest{})
	if err != nil {
		errStatus, _ := status.FromError(err)
		fmt.Println(errStatus.Message())
		fmt.Println(errStatus.Code())

		for _, d := range errStatus.Details() {
			switch info := d.(type) {
			default:
				log.Printf("Details type: %s", info)
			}
		}

		return nil, fmt.Errorf("wrapped: %w", ErrGetMinerStats{err.Error()})
	}

	return minerStatsResponse, nil
}

// Get Miner Firmware Version
func GetAPIVersion(authCtx context.Context) (*pb.ApiVersion, error) {

	apiVersion, err := apiVersionClient.GetApiVersion(authCtx, &pb.ApiVersionRequest{})
	if err != nil {
		return nil, err
	}

	return apiVersion, nil
}

// LoginWrapper
func (app *Config) LoginWrapper() error {
	headerMD := metadata.MD{}

	loginResponse, err := Login(&headerMD)
	if err != nil {
		log.Fatalf("could not login: %v", err)
		return err
	} else {
		log.Println("Login successful")
	}

	authTokens := headerMD.Get(app.AuthToken.Key)
	if len(authTokens) == 0 {
		log.Fatal("authorization token not found in headers")
	}
	app.AuthToken.Value = authTokens[0] // Taking the first token
	app.AuthToken.TimeOutS = int(loginResponse.TimeoutS)

	log.Println("authTokenValue: ", app.AuthToken.Value)

	// Attach auth token to context
	md := metadata.New(map[string]string{"authorization": app.AuthToken.Value})
	authCtx = metadata.NewOutgoingContext(ctx, md)

	return nil
}

// Login
func Login(headerMD *metadata.MD) (*pbV1.LoginResponse, error) {
	loginReq := &pbV1.LoginRequest{
		Username: app.MinerUsername,
		Password: app.MinerPassword,
	}

	loginResponse, err := authClient.Login(ctx, loginReq, grpc.Header(headerMD))
	if err != nil {
		return nil, err
	}

	return loginResponse, nil
}

// Stops Miner
func MinerStop(authCtx context.Context) (interface{}, error) {
	_, err := actionsClient.Stop(authCtx, &pbV1.StopRequest{})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Starts Miner
func MinerStart(authCtx context.Context) (interface{}, error) {
	_, err := actionsClient.Start(authCtx, &pbV1.StartRequest{})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Set Power Target of the Miner
func MinerSetPowerTarget(authCtx context.Context, powerThreshold uint64) (interface{}, error) {
	_, err := performanceClient.SetPowerTarget(authCtx, &pbV1.SetPowerTargetRequest{
		SaveAction: pbV1.SaveAction_SAVE_ACTION_SAVE_AND_APPLY,
		PowerTarget: &pbV1.Power{
			Watt: powerThreshold,
		},
	})

	if err != nil {
		return nil, err
	}

	return nil, nil
}

package main

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

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
func GetMinerConfiguration(args ...interface{}) (interface{}, error) {
	authCtx := args[0].(context.Context)
	minerConfigResponse, err := configClient.GetMinerConfiguration(authCtx, &pbV1.GetMinerConfigurationRequest{})
	if err != nil {
		return nil, err
	}

	return minerConfigResponse, nil
}

// Get Miner Firmware Version
func GetAPIVersion(args ...interface{}) (interface{}, error) {
	authCtx := args[0].(context.Context)
	apiVersion, err := apiVersionClient.GetApiVersion(authCtx, &pb.ApiVersionRequest{})
	if err != nil {
		return nil, err
	}

	return apiVersion, nil
}

// Login
func Login(args ...interface{}) (interface{}, error) {
	headerMD := args[0].(*metadata.MD)

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
func MinerStop(args ...interface{}) (interface{}, error) {
	authCtx := args[0].(context.Context)
	_, err := actionsClient.Stop(authCtx, &pbV1.StopRequest{})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Starts Miner
func MinerStart(args ...interface{}) (interface{}, error) {
	authCtx := args[0].(context.Context)
	_, err := actionsClient.Start(authCtx, &pbV1.StartRequest{})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Set Power Target of the Miner
func MinerSetPowerTarget(args ...interface{}) (interface{}, error) {
	authCtx := args[0].(context.Context)
	powerThreshold := args[1].(uint64)

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

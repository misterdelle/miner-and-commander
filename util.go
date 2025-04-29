package main

import (
	"encoding/json"
	pbV1 "github.com/misterdelle/miner-and-commander/pb/github.com/braiins/bos-plus-api/braiins/bos/v1"
	"github.com/misterdelle/miner-and-commander/pkg/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"strings"
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
	app.MinerOperations.AuthCtx = authCtx

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

func (app *Config) faiLaMediaDelleLetture() model.Station {
	ms := model.NewStation()
	var currentTotalPowerFromPV float32 = 0.0
	var currentConsumptionPower float32 = 0.0
	var currentGridFeedPower float32 = 0.0
	var currentBatteryPower float32 = 0.0

	it := pvRead.Iterator()
	for it.Next() {
		_, value := it.Index(), it.Value().(model.Station)
		//log.Printf("index: %v, value: %v", index, value.LastUpdateTime)
		ms.LastUpdateTime = value.LastUpdateTime
		ms.LastUpdateTimeRead = true
		ms.CurrentBatterySOC = value.CurrentBatterySOC
		ms.CurrentBatterySOCRead = true
		currentTotalPowerFromPV += value.CurrentTotalPowerFromPV
		currentConsumptionPower += value.CurrentConsumptionPower
		currentGridFeedPower += value.CurrentGridFeedPower
		currentBatteryPower += value.CurrentBatteryPower
	}

	ms.CurrentTotalPowerFromPV = currentTotalPowerFromPV / float32(pvRead.Size())
	ms.CurrentTotalPowerFromPVRead = true
	ms.CurrentConsumptionPower = currentConsumptionPower / float32(pvRead.Size())
	ms.CurrentConsumptionPowerRead = true
	ms.CurrentGridFeedPower = currentGridFeedPower / float32(pvRead.Size())
	ms.CurrentGridFeedPowerRead = true
	ms.CurrentBatteryPower = currentBatteryPower / float32(pvRead.Size())
	ms.CurrentBatteryPowerRead = true

	return ms
}

package main

import (
	"encoding/json"
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

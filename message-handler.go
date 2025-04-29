package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	//log.Printf("Received message: %s from topic: %s\n", message.Payload(), message.Topic())

	parseStationData(message.Payload(), message.Topic())
}

func parseStationData(msgPayload []byte, msgTopic string) {

	if msgTopic == app.MQTTTopicName+"/station/totalPowerFromPV" {
		stationData.CurrentTotalPowerFromPV = fromByteArrayToFloat32(msgPayload)
		stationData.CurrentTotalPowerFromPVRead = true
	}

	if msgTopic == app.MQTTTopicName+"/station/currentConsumptionPower" {
		stationData.CurrentConsumptionPower = fromByteArrayToFloat32(msgPayload)
		stationData.CurrentConsumptionPowerRead = true
	}

	if msgTopic == app.MQTTTopicName+"/LoadInfo/Load C/Load C Power" {
		stationData.CurrentGridFeedPower = fromByteArrayToFloat32(msgPayload)
		stationData.CurrentGridFeedPowerRead = true
	}

	if msgTopic == app.MQTTTopicName+"/station/batteryPower" {
		stationData.CurrentBatteryPower = fromByteArrayToFloat32(msgPayload)
		stationData.CurrentBatteryPowerRead = true
	}

	if msgTopic == app.MQTTTopicName+"/station/batterySOC" {
		stationData.CurrentBatterySOC = fromByteArrayToFloat32(msgPayload)
		stationData.CurrentBatterySOCRead = true
	}

	if msgTopic == app.MQTTTopicName+"/station/lastUpdateTime" {
		// 2024-01-19 10:28:48
		lastUpdate := strings.TrimSpace(string(msgPayload))
		lastUpdateTS, err := time.Parse("2006-01-02 15:04:05", lastUpdate)
		if err != nil {
			log.Printf("Error parsing lastUpdateTime: %s", err)
		}
		if stationData.LastUpdateTime != lastUpdateTS {
			//log.Printf("Old Timestamp: %s\n", stationData.LastUpdateTime)
			//log.Printf("New Timestamp: %s\n", lastUpdateTS)
			//log.Println("Invalidating PV Station Data")
			stationData.Invalidate()
			//log.Println("Updating PV Station Data")
			stationData.LastUpdateTime = lastUpdateTS
			stationData.LastUpdateTimeRead = true
		}
		//log.Printf("Received message: %s from topic: %s\n", lastUpdate, msgTopic)
	}

	if stationData.IsValid() {
		//log.Printf("PV Station Data : %s \n", stationData.ToJSON())
		pvRead.Enqueue(stationData)
		stationData.Invalidate()
	}
}

func fromByteArrayToFloat32(b []byte) float32 {
	num, err := strconv.ParseFloat(strings.TrimSpace(string(b)), 64)
	if err != nil {
		fmt.Println("fromByteArrayToFloat32 failed:", err)

		return 0
	}

	return float32(num)
}

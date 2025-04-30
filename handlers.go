package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type ErrGetMinerStats struct {
	s string
}

func (app *Config) GetMinersDetails(w http.ResponseWriter, r *http.Request) {
	minerDetailsResponse, err := app.MinerOperations.GetMinerDetails()
	if err != nil {
		log.Fatalf("could not get miner details: %v", err)
	}

	jsonMinerDetailsResponse, err := json.Marshal(minerDetailsResponse)
	if err != nil {
		log.Fatalf("could not convert to json miner details: %v", err)
	}
	log.Println("jsonMinerDetailsResponse: ", string(jsonMinerDetailsResponse))

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(jsonMinerDetailsResponse))
}

func (app *Config) GetMinersStats(w http.ResponseWriter, r *http.Request) {
	var myErr = ErrGetMinerStats{}

	minerStatsResponse, err := app.MinerOperations.GetMinerStats()
	if err != nil {
		// if err.Error() == "BOSminer API connection error: Connection refused (os error 111)" {
		if errors.As(err, &myErr) {
			log.Printf("ErrGetMinerStats: %v\n", err)

			if myErr.s == "rpc error: code = Internal desc = BOSminer API connection error: Connection refused (os error 111)" {
				log.Println("GetMinersStats ERROR: BOSminer API connection error: Connection refused")
			} else if myErr.s == "rpc error: code = FailedPrecondition desc = BOSminer is not running" {
				log.Println("GetMinersStats ERROR: BOSminer is not running")
			}
		} else {
			log.Printf("could not get miner stats: %v\n", err)
		}
	}

	jsonMinerStatsResponse, err := json.Marshal(minerStatsResponse)
	if err != nil {
		log.Fatalf("could not convert to json miner stats: %v", err)
	}
	log.Println("jsonMinerStatsResponse: ", string(jsonMinerStatsResponse))

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(jsonMinerStatsResponse))
}

func (app *Config) GetMinerConfiguration(w http.ResponseWriter, r *http.Request) {
	minerConfigResponse, err := app.MinerOperations.GetMinerConfiguration()
	if err != nil {
		log.Fatalf("could not get miner configuration: %v", err)
	}

	jsonMinerConfigResponse, err := json.Marshal(minerConfigResponse)
	if err != nil {
		log.Fatalf("could not convert to json miner configuration: %v", err)
	}
	log.Println("jsonMinerConfigResponse: ", string(jsonMinerConfigResponse))

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(jsonMinerConfigResponse))
}

func (app *Config) GetPVData(w http.ResponseWriter, r *http.Request) {
	pvDataResponse, err := json.Marshal(stationData)
	if err != nil {
		log.Fatalf("could not convert to json station data: %v", err)
	}
	log.Println("pvDataResponse: ", string(pvDataResponse))
	stationData = app.faiLaMediaDelleLetture()
	pvDataResponse, err = json.Marshal(stationData)
	if err != nil {
		log.Fatalf("could not convert to json station data: %v", err)
	}
	log.Println("pvDataResponse: ", string(pvDataResponse))

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(pvDataResponse))
}

func (app *Config) DoCheck(w http.ResponseWriter, r *http.Request) {
	app.startCheck()

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("OK"))
}

func (app *Config) ApplyConfig(w http.ResponseWriter, r *http.Request) {
	cfgName := r.URL.Query().Get("cfg-name")
	if cfgName != "" {
		app.applyConfig(cfgName)
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("OK"))
}

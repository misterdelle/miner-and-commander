package model

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Station struct {
	LastUpdateTime              time.Time
	LastUpdateTimeRead          bool    `json:"-"`
	CurrentTotalPowerFromPV     float32 // gosolarmanpv/station/totalPowerFromPV
	CurrentTotalPowerFromPVRead bool    `json:"-"`
	CurrentConsumptionPower     float32 // gosolarmanpv/station/currentConsumptionPower
	CurrentConsumptionPowerRead bool    `json:"-"`
	CurrentGridFeedPower        float32 // gosolarmanpv/LoadInfo/Load C/Load C Power
	CurrentGridFeedPowerRead    bool    `json:"-"`
	CurrentBatteryPower         float32 // gosolarmanpv/station/batteryPower
	CurrentBatteryPowerRead     bool    `json:"-"`
	CurrentBatterySOC           float32 // gosolarmanpv/station/batterySOC
	CurrentBatterySOCRead       bool    `json:"-"`
}

func NewStation() Station {
	return Station{}
}

func (s Station) String() string {
	ret := "Station: \n"
	ret += fmt.Sprintf("LastUpdateTime: %s  \n", s.LastUpdateTime.String())
	ret += fmt.Sprintf("TotalPowerFromPV: %.2f \n", s.CurrentTotalPowerFromPV)
	ret += fmt.Sprintf("ConsumptionPower: %.2f \n", s.CurrentConsumptionPower)
	ret += fmt.Sprintf("GridFeedPower: %.2f \n", s.CurrentGridFeedPower)
	ret += fmt.Sprintf("BatteryPower: %.2f \n", s.CurrentBatteryPower)
	ret += fmt.Sprintf("BatterySOC: %.2f \n", s.CurrentBatterySOC)

	return ret
}

func (s Station) IsValid() bool {
	return s.LastUpdateTimeRead && s.CurrentBatteryPowerRead && s.CurrentGridFeedPowerRead && s.CurrentConsumptionPowerRead && s.CurrentTotalPowerFromPVRead && s.CurrentBatterySOCRead
}

func (s Station) Invalidate() {
	s.LastUpdateTimeRead = false
	s.CurrentBatteryPowerRead = false
	s.CurrentGridFeedPowerRead = false
	s.CurrentConsumptionPowerRead = false
	s.CurrentTotalPowerFromPVRead = false
	s.CurrentBatterySOCRead = false
}

func (s Station) ToJSON() string {
	retJSON, err := json.Marshal(s)
	if err != nil {
		log.Fatalf("could not convert to json station data: %v", err)
	}

	return string(retJSON)
}

package model

import (
	"strconv"
	"strings"
)

const HASHBOARD_NUMBER = 3

type MinerThreshold struct {
	MinPowerThreshold      uint64
	MaxPowerThreshold      uint64
	MinerConfigurationName string
}

type MinerConfiguration struct {
	Name           string
	PowerThreshold uint64
	HashboardIds   []string
}

func LoadMinerConfigurationsMap(minerConfigurations map[string]MinerConfiguration) {
	minerConfigurations["0"] = MinerConfiguration{
		Name:           "0W",
		PowerThreshold: 0,
		HashboardIds:   []string{""},
	}

	minerConfigurations["300"] = MinerConfiguration{
		Name:           "300W",
		PowerThreshold: 966,
		HashboardIds:   []string{"1"},
	}

	minerConfigurations["600"] = MinerConfiguration{
		Name:           "600W",
		PowerThreshold: 966,
		HashboardIds:   []string{"1", "2"},
	}

	minerConfigurations["966"] = MinerConfiguration{
		Name:           "966W",
		PowerThreshold: 966,
		HashboardIds:   []string{"1", "2", "3"},
	}

	minerConfigurations["1200"] = MinerConfiguration{
		Name:           "1200W",
		PowerThreshold: 1200,
		HashboardIds:   []string{"1", "2", "3"},
	}

	minerConfigurations["1500"] = MinerConfiguration{
		Name:           "1500W",
		PowerThreshold: 1500,
		HashboardIds:   []string{"1", "2", "3"},
	}

	minerConfigurations["2000"] = MinerConfiguration{
		Name:           "2000W",
		PowerThreshold: 2000,
		HashboardIds:   []string{"1", "2", "3"},
	}

	minerConfigurations["2500"] = MinerConfiguration{
		Name:           "2500W",
		PowerThreshold: 2500,
		HashboardIds:   []string{"1", "2", "3"},
	}

	minerConfigurations["3068"] = MinerConfiguration{
		Name:           "3068W",
		PowerThreshold: 3068,
		HashboardIds:   []string{"1", "2", "3"},
	}

}

func ParseMinerThreshold(line string) *MinerThreshold {
	s := strings.Split(line, ",")
	minPowerThreshold, _ := strconv.Atoi(s[0])
	maxPowerThreshold, _ := strconv.Atoi(s[1])
	minerConfigurationName := s[2]
	return &MinerThreshold{MinPowerThreshold: uint64(minPowerThreshold), MaxPowerThreshold: uint64(maxPowerThreshold), MinerConfigurationName: minerConfigurationName}
}

func GetMinerConfigurationNameByThreshold(minerThresholdList []*MinerThreshold, threshold uint64) string {
	rc := ""
	for i := range minerThresholdList {
		x := minerThresholdList[i]
		if x.MinPowerThreshold <= threshold && threshold <= x.MaxPowerThreshold {
			rc = x.MinerConfigurationName
			break
		}
	}

	return rc
}

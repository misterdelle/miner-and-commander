package model

import (
	"strconv"
	"strings"
)

type MinerThreshold struct {
	MinPowerThreshold      uint64
	MaxPowerThreshold      uint64
	MinerConfigurationName string
}

func ParseMinerThreshold(line string) *MinerThreshold {
	s := strings.Split(line, ",")
	minPowerThreshold, _ := strconv.Atoi(s[0])
	maxPowerThreshold, _ := strconv.Atoi(s[1])
	minerConfigurationName := s[2]
	return &MinerThreshold{MinPowerThreshold: uint64(minPowerThreshold), MaxPowerThreshold: uint64(maxPowerThreshold), MinerConfigurationName: minerConfigurationName}
}

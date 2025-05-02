package model

const HASHBOARD_NUMBER = 3

type MinerConfiguration struct {
	Name           string
	PowerThreshold uint64
	HashboardIds   []string
}

func NewMinerConfiguration() MinerConfiguration {
	return MinerConfiguration{}
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

func GetMinerConfigurationByThreshold(minerConfigurations map[string]MinerConfiguration, minerThresholdList []*MinerThreshold, threshold uint64) MinerConfiguration {
	rc := NewMinerConfiguration()

	for i := range minerThresholdList {
		x := minerThresholdList[i]
		if x.MinPowerThreshold <= threshold && threshold <= x.MaxPowerThreshold {
			rc = minerConfigurations[x.MinerConfigurationName]
			break
		}
	}

	return rc
}

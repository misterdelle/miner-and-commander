package miner_ops

import (
	"context"
	"fmt"
	pb "github.com/misterdelle/miner-and-commander/pb/github.com/braiins/bos-plus-api/braiins/bos"
	pbV1 "github.com/misterdelle/miner-and-commander/pb/github.com/braiins/bos-plus-api/braiins/bos/v1"
	"github.com/misterdelle/miner-and-commander/pkg/model"
	"google.golang.org/grpc/status"
	"log"
	"math/rand/v2"
	"time"
)

type ErrGetMinerStats struct {
	s string
}

type MinerOps struct {
	AuthCtx            context.Context
	configClient       pbV1.ConfigurationServiceClient
	actionsClient      pbV1.ActionsServiceClient
	minerServiceClient pbV1.MinerServiceClient
	performanceClient  pbV1.PerformanceServiceClient
	apiVersionClient   pb.ApiVersionServiceClient
}

func NewMinerOps(authCtx context.Context, apiVersionClient pb.ApiVersionServiceClient, configClient pbV1.ConfigurationServiceClient, actionsClient pbV1.ActionsServiceClient, minerServiceClient pbV1.MinerServiceClient, performanceClient pbV1.PerformanceServiceClient) *MinerOps {
	return &MinerOps{
		AuthCtx:            authCtx,
		apiVersionClient:   apiVersionClient,
		configClient:       configClient,
		actionsClient:      actionsClient,
		minerServiceClient: minerServiceClient,
		performanceClient:  performanceClient,
	}
}

// Get Miner Configuration
func (mo *MinerOps) GetMinerConfiguration() (*pbV1.GetMinerConfigurationResponse, error) {
	minerConfigResponse, err := mo.configClient.GetMinerConfiguration(mo.AuthCtx, &pbV1.GetMinerConfigurationRequest{})
	if err != nil {
		return nil, err
	}

	return minerConfigResponse, nil
}

func (mo *MinerOps) SetMinerConfiguration(currentMinerConfig, newMinerConfig *model.MinerConfiguration) {
	if newMinerConfig.Name == "0W" {
		//
		// Se la powerThreshold è uguale zero spengo il miner
		//
		log.Println("Stopping miner")

		_, err := mo.MinerStop()
		if err != nil {
			log.Println("could not stop miner", err)
		}

		return
	}

	if len(newMinerConfig.HashboardIds) == 1 {
		var randId = rand.IntN(model.HASHBOARD_NUMBER)
		if randId == 0 {
			randId++
		}
		newMinerConfig.HashboardIds = []string{fmt.Sprintf("%d", randId)}
	} else if len(newMinerConfig.HashboardIds) == 2 {
		var randId = rand.IntN(model.HASHBOARD_NUMBER)
		if randId == 0 {
			randId++
		}
		id := fmt.Sprintf("%d", randId)
		var allHashboardIds = getAllHashboardIds()
		var newHashboardIds []string
		for v := range allHashboardIds {
			appoId := allHashboardIds[v]
			if appoId != id {
				newHashboardIds = append(newHashboardIds, appoId)
			}
		}
		newMinerConfig.HashboardIds = newHashboardIds
	}

	//
	// Nel dubbio faccio partire il miner, poi setterò i dettagli
	//
	log.Println("Starting miner")

	_, err := mo.MinerStart()
	if err != nil {
		log.Println("could not start miner", err)
	}

	//
	// Setto il miner con la configurazione prevista
	//
	log.Printf("Setting Power Target to %v with Hashboards: %s\n", newMinerConfig.PowerThreshold, newMinerConfig.HashboardIds)

	if currentMinerConfig.PowerThreshold == newMinerConfig.PowerThreshold {
		//
		// Se la configurazione corrente ha la stessa powerThreshold riavvio il miner
		//
		_, err = mo.MinerSetHashboardsAndPowerTarget(*newMinerConfig, true)
		if err != nil {
			log.Printf("could not set power target: %v", err)
		}
	} else {
		_, err = mo.MinerSetHashboardsAndPowerTarget(*newMinerConfig, false)
		if err != nil {
			log.Printf("could not set power target: %v", err)
		}
	}
}

// Get Miner Details
func (mo *MinerOps) GetMinerDetails() (*pbV1.GetMinerDetailsResponse, error) {
	minerDetailsResponse, err := mo.minerServiceClient.GetMinerDetails(mo.AuthCtx, &pbV1.GetMinerDetailsRequest{})
	if err != nil {
		return nil, err
	}

	return minerDetailsResponse, nil
}

// Get Miner Stats
func (mo *MinerOps) GetMinerStats() (*pbV1.GetMinerStatsResponse, error) {
	minerStatsResponse, err := mo.minerServiceClient.GetMinerStats(mo.AuthCtx, &pbV1.GetMinerStatsRequest{})
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
func (mo *MinerOps) GetAPIVersion() (*pb.ApiVersion, error) {

	apiVersion, err := mo.apiVersionClient.GetApiVersion(mo.AuthCtx, &pb.ApiVersionRequest{})
	if err != nil {
		return nil, err
	}

	return apiVersion, nil
}

// Stops Miner
func (mo *MinerOps) MinerStop() (interface{}, error) {
	_, err := mo.actionsClient.Stop(mo.AuthCtx, &pbV1.StopRequest{})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Starts Miner
func (mo *MinerOps) MinerStart() (interface{}, error) {
	_, err := mo.actionsClient.Start(mo.AuthCtx, &pbV1.StartRequest{})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Set Power Target of the Miner
func (mo *MinerOps) MinerSetPowerTarget(powerThreshold uint64) (interface{}, error) {
	_, err := mo.performanceClient.SetPowerTarget(mo.AuthCtx, &pbV1.SetPowerTargetRequest{
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

// Set Hashboards Enabled and the Power Target of the Miner
func (mo *MinerOps) MinerSetHashboardsAndPowerTarget(mc model.MinerConfiguration, restartMiner bool) (interface{}, error) {
	allHashboardIds := getAllHashboardIds()
	var powerThreshold uint64 = mc.PowerThreshold
	var hashboardIds []string = mc.HashboardIds

	//
	// Disabilito tutte le Hashboards
	//
	_, err := mo.minerServiceClient.DisableHashboards(mo.AuthCtx, &pbV1.DisableHashboardsRequest{
		SaveAction:   pbV1.SaveAction_SAVE_ACTION_SAVE,
		HashboardIds: allHashboardIds,
	})
	if err != nil {
		return nil, err
	}

	//
	// Abilito le Hashboards necessarie
	//
	_, err = mo.minerServiceClient.EnableHashboards(mo.AuthCtx, &pbV1.EnableHashboardsRequest{
		SaveAction:   pbV1.SaveAction_SAVE_ACTION_SAVE,
		HashboardIds: hashboardIds,
	})
	if err != nil {
		return nil, err
	}

	//
	// Setto la Powerthreshold e applico le modifiche
	//
	_, err = mo.performanceClient.SetPowerTarget(mo.AuthCtx, &pbV1.SetPowerTargetRequest{
		SaveAction: pbV1.SaveAction_SAVE_ACTION_SAVE_AND_APPLY,
		PowerTarget: &pbV1.Power{
			Watt: powerThreshold,
		},
	})

	if restartMiner {
		log.Println("Stopping miner")

		_, err := mo.MinerStop()
		if err != nil {
			log.Println("could not stop miner", err)
		}

		time.Sleep(time.Second * 30)

		log.Println("Starting miner")

		_, err = mo.MinerStart()
		if err != nil {
			log.Println("could not start miner", err)
		}
	}

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func getAllHashboardIds() []string {
	var hashboardIds []string
	for i := 1; i <= model.HASHBOARD_NUMBER; i++ {
		hashboardIds = append(hashboardIds, fmt.Sprintf("%d", i))
	}
	return hashboardIds
}

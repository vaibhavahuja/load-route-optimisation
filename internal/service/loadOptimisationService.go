package service

import (
	log "github.com/sirupsen/logrus"
	"github.com/vaibhavahuja/load-route-optimisation/internal/entities"
)

func (app *Application) CalculateOptimalLotSize(optimalLotSizeRequest entities.OptimalLotSizeRequest) (optimalLotSizeResponse entities.OptimalLotSizeResponse, err error) {
	log.Infof("received the following request %v", optimalLotSizeRequest)
	//sample output check
	optimalLotSizeResponse.ItemsToBeAdded = optimalLotSizeRequest.Items
	return
}

//This method contains the implementation of knapsack
//Currently in progress
func knapsack() {

}

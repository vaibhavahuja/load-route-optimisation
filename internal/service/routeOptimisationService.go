package service

import (
	log "github.com/sirupsen/logrus"
	"github.com/vaibhavahuja/load-route-optimisation/internal/entities"
	"github.com/vaibhavahuja/load-route-optimisation/utils"
)

func (app *Application) FindOptimalRoute(request entities.OptimalRouteRequest) (response entities.OptimalRouteResponse, err error) {
	log.Infof("received the following request %v", request)

	startNode := request.StartLocation.GetNodeFromLocation()
	endNode := request.EndLocation.GetNodeFromLocation()
	//startNode := &entities.Node{Value: "A"}
	//endNode := &entities.Node{Value: "B"}
	log.Infof("start node is %v and end node is %v", startNode, endNode)

	list, val := utils.GetShortestPath(startNode, endNode, app.graph)

	log.Info(val)
	response.Value = list
	return
}

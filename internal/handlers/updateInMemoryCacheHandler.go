package handlers

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/vaibhavahuja/load-route-optimisation/internal/entities"
	"net/http"
)

func (server *HttpServer) UpdateInMemoryCache(resp http.ResponseWriter, req *http.Request) {
	log.Info("Received request to update the value of cache ", req)
	log.Info("received request")
	var request entities.UpdateCacheRequest
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&request)
	cacheFinalResponse, err := server.svc.UpdateValueInMemoryCache(request)
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(cacheFinalResponse)
	if err != nil {
		return
	}
	_, err = resp.Write(jsonResponse)
	if err != nil {
		log.Errorf("Unable to write error %s", err)
	}
	log.Info("Successfully published record to API")

}

func (server *HttpServer) MultsiNodeRouteOptimisationHandler(resp http.ResponseWriter, req *http.Request) {
	log.Info("received request")
	var request entities.OptimalRouteMultipleNodesRequest
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&request)
	if err != nil {
		log.Errorf("Error while decoding request %s", err)
	}
	log.Info("sending output to request")
	response, err := server.svc.FindOptimalRouteAcrossMultiNodes(request)
	if err != nil {
		log.Errorf("Erorr while calculating optimal route : %s", err)
		resp.WriteHeader(http.StatusInternalServerError)
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	_, err = resp.Write(jsonResponse)
	if err != nil {
		log.Errorf("Unable to write error %s", err)
	}
	log.Info("Successfully published record to API")
}

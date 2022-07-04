package main

import (
	"context"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/vaibhavahuja/load-route-optimisation/internal/handlers"
	"github.com/vaibhavahuja/load-route-optimisation/internal/service"
	"github.com/vaibhavahuja/load-route-optimisation/utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Info("Starting Http Server")
	//create new graph here and inject it

	//will be calling a script instead to get this data from json instead

	//for local testing I am injecting my own graph
	myGraphInputData := []utils.InputData{
		{
			Source:      "lat28.641529#long77.120918",
			Destination: "lat28.699774#long77.138596",
			Weight:      10,
		},
		{
			Source:      "lat28.641529#long77.120918",
			Destination: "C",
			Weight:      50,
		},
		{
			Source:      "C",
			Destination: "D",
			Weight:      30,
		},
		{
			Source:      "D",
			Destination: "lat28.699774#long77.138596",
			Weight:      10,
		},
	}

	inputGraph := utils.InputGraph{
		Graph: myGraphInputData,
	}
	graphCreated := utils.CreateGraph(inputGraph)

	log.Infof("successfully initialised graph with value %v", graphCreated)

	app := service.NewApplication(graphCreated)
	svc := handlers.NewHttpServer(app)
	//creating a new router
	r := mux.NewRouter()

	//Registering all handlers
	r.HandleFunc("/health-check", svc.HealthCheck).Methods("GET")
	r.HandleFunc("/api/v1/optimise-load", svc.LoadOptimisationHandler).Methods("POST")
	r.HandleFunc("/api/v1/optimise-route", svc.RouteOptimisationHandler).Methods("POST")
	r.HandleFunc("/api/v1/report-blocker", svc.ReportRouteBlockerHandler).Methods("POST")

	server := RunHttpServer(":8080", r)
	log.Info("Successfully started server")
	gracefulStop(server)
}

func RunHttpServer(port string, router *mux.Router) *http.Server {
	httpServer := &http.Server{
		Addr:    port,
		Handler: router,
	}

	go func() {
		log.Infof("Starting server on localhost%v", port)
		if err := httpServer.ListenAndServe(); err != nil {
			log.Errorf("Error in creating server : %s", err)
		}
	}()
	return httpServer
}

func gracefulStop(gs *http.Server) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	signal := <-c
	log.Infof("Stopping the server. signal : %s", signal)
	ctx, cancel := context.WithTimeout(context.Background(), 2000)
	defer cancel()
	if err := gs.Shutdown(ctx); err != nil {
		log.Errorf("failed to shutdown correctly. Error: %s", err)
	}
}

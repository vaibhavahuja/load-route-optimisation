package service

import "github.com/vaibhavahuja/load-route-optimisation/utils"

type Application struct {
	graph   *utils.ItemGraph
	myCache string
}

func NewApplication(itemGraph *utils.ItemGraph, text string) *Application {
	return &Application{
		graph:   itemGraph,
		myCache: text,
	}
}

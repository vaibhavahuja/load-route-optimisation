package main

import (
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Run(address string) {

}

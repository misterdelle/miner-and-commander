package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Config) routes() http.Handler {
	// create router
	mux := chi.NewRouter()

	// set up middleware
	mux.Use(middleware.Recoverer)
	// Register with router
	mux.Use(Logger)
	// Register with router
	// mux.Use(AuthChecker)

	// define application routes
	mux.Get("/MinerConfiguration", app.GetMinerConfiguration)
	mux.Get("/MinerDetails", app.GetMinersDetails)
	mux.Get("/MinerStats", app.GetMinersStats)
	mux.Get("/PVData", app.GetPVData)
	mux.Get("/DoCheck", app.DoCheck)

	return mux
}

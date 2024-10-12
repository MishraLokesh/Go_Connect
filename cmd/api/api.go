package main

import (
	"log"
	"net/http"
	"time"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/health", app.healthCheckHandler)

	return mux
}

func (app application) run(mux *http.ServeMux) error {
	s := &http.Server{
		Addr: app.config.addr,
		Handler: mux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Server has started at %s", app.config.addr)

	return s.ListenAndServe()
}



package main

import (
	"log"
	"net/http"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app application) run() error {
	mux := http.NewServeMux()

	s := &http.Server{
		Addr: app.config.addr,
		Handler: mux,
	}

	log.Printf("Server has started at %s", app.config.addr)

	return s.ListenAndServe()
}



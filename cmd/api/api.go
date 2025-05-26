package main

import (
	"log"
	"net/http"
	"rest-api/internal/store"
	"time"
)

type application struct {
	config  config
	store   store.Storage
	infoLog *log.Logger
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type config struct {
	addr        string
	db          dbConfig
	env         string
	apiURL      string
	frontendURL string
}

func (app *application) run(mux http.Handler) error {

	srv := &http.Server{
		Addr:         ":"+ app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	err := srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
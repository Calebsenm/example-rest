package main

import (

	httpSwagger "github.com/swaggo/http-swagger/v2"
	"net/http"

)

func (app *application) routes() http.Handler {

	mux := http.NewServeMux()

	mux.Handle("/docs/", httpSwagger.WrapHandler)

	mux.HandleFunc("POST /api/project/create", app.createProject)
	mux.HandleFunc("GET /api/project/get_alls", app.getallProjects)

	mux.HandleFunc("POST /api/participant/create", app.createParticipant)
	mux.HandleFunc("GET /api/participant/get_alls", app.getallParticipant)

	return  app.enableCORS(app.logRequest(mux))
}

package main

import (
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"net/http"
)

func (app *application) routes() http.Handler {

	mux := http.NewServeMux()

	mux.Handle("/docs/", httpSwagger.WrapHandler)

	mux.HandleFunc("POST /api/projects", app.createProject)
	mux.HandleFunc("GET /api/projects", app.getallProjects)
	mux.HandleFunc("GET /api/project", app.searchProject)


	mux.HandleFunc("POST /api/participants", app.createParticipant)
	mux.HandleFunc("GET /api/participants", app.getallParticipant)

	mux.HandleFunc("POST /api/assignments", app.assigmentProject)
	mux.HandleFunc("GET /api/assignments", app.getAssignments)

	return app.enableCORS(app.logRequest(mux))
}

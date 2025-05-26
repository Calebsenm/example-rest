package main

import (
	"net/http"
	"rest-api/internal/store"
)

// @Summary Crear participante
// @Description Crea un nuevo proyecto
// @Tags Proyectos
// @Accept json
// @Produce json
// @Param participant body store.Project true "Datos del participante"
// @Success 201 {object} store.Project
// @Router /project/create [post]
func (app *application) createProject(w http.ResponseWriter, r *http.Request) {

	var input store.Project

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	project := &store.Project{
		ProjectID:   input.ProjectID,
		Name:        input.Name,
		Description: input.Description,
		StartDate:   input.StartDate,
		EndDate:     input.EndDate,
		Value:       input.Value,
	}

	ctx := r.Context()

	err = app.store.Projects.Create(ctx, project)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"project": project}, nil)

	if err != nil {
		app.internalServerError(w, r, err)
	}
}


// getallParticipant obtiene todos los proyectos
// @Summary Obtener todos los proyectos
// @Description Retorna la lista de todos los proyectos registrados
// @Tags Proyectos
// @Produce json
// @Success 200 {object} []store.Project "Lista de proyectos"
// @Router /project/get_alls [get]
func (app *application) getallProjects(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	projects , err := app.store.Projects.GetAlls(ctx)
	if err != nil{
		app.internalServerError(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"proyects": projects}, nil)

	if err != nil {
		app.internalServerError(w, r, err)
	}
}

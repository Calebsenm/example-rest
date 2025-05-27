package main

import (
	"net/http"
	"rest-api/internal/store"
)
// @Summary Asignar Proyecto 
// @Description Asigna un proyecto a un estudiante
// @Tags AsignarProyecto 
// @Accept json
// @Produce json
// @Param assignment body store.Assignment true "Id para asignar un proyecto"
// @Success 201 {object} store.Assignment
// @Router /assignments [post]
func (app *application) assigmentProject(w http.ResponseWriter , r *http.Request){
	var input  store.Assignment 
	
	err := app.readJSON(w , r ,&input)
	if err != nil {
		app.internalServerError(w , r , err)
		return 
	}

	assignament := &store.Assignment{
		ProjectID:  input.ProjectID,
		ParticipantID: input.ParticipantID,
	}

	ctx := r.Context()

	err = app.store.Assignment.AssignProject(ctx , assignament)
	if err != nil {
		app.internalServerError(w , r , err)
		return 
	}

	err = app.writeJSON(w , http.StatusCreated , envelope{"assignament":assignament} , nil)
	if err != nil { 
		app.internalServerError(w , r , err)
	}
}
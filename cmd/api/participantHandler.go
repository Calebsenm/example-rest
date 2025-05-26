package main

import (
	"net/http"
	"rest-api/internal/store"
)

// @Summary Crear participante
// @Description Crea un nuevo participante 
// @Tags Participantes
// @Accept json
// @Produce json
// @Param participant body store.Participants true "Datos del participante"
// @Success 201 {object} store.Participants
// @Router /participant/create [post]
func (app *application) createParticipant(w http.ResponseWriter, r *http.Request) {

	var input store.Participants

	if err := app.readJSON(w, r, &input); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	participant := &store.Participants{
		Identification: input.Identification,
		First_name:     input.First_name,
		Last_name:      input.First_name,
		Email:          input.Email,
		Phone:          input.Phone,
	}

	ctx := r.Context()
	err := app.store.Participants.Create(ctx, participant)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"participants": participant}, nil)

	if err != nil {
		app.internalServerError(w, r, err)
	}
}


// getallParticipant obtiene todos los participantes
// @Summary Obtener todos los participantes
// @Description Retorna la lista de todos los participantes registrados
// @Tags Participantes
// @Produce json
// @Success 200 {object} []store.Participants "Lista de participantes"
// @Router /participant/get_alls [get]
func (app *application) getallParticipant(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	participants , err := app.store.Participants.GetAlls(ctx)

	if err != nil{
		app.internalServerError(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"participants": participants }, nil)

	if err != nil {
		app.internalServerError(w, r, err)
	}
}

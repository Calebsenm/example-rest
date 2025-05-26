package main

import "net/http"

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {

	app.infoLog.Printf("[ERROR] internal | method=%s | path=%s | error=%s",
		r.Method, r.URL.Path, err.Error())
		
	app.writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}

package main

import (
	"net/http"
)

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)

	})
}


func (app *application) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Permite cualquier origen
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Permite los métodos que desees
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Permite headers personalizados
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Responder a las solicitudes OPTIONS sin procesar más
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Continua con el siguiente handler
		next.ServeHTTP(w, r)
	})
}
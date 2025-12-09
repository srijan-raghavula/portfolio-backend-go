package main

import (
	"encoding/json"
	"net/http"
)

type Payload struct {
	Message string `json:"message"`
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) sendMessage(payload *Payload, w http.ResponseWriter, status int) {
	body, err := json.Marshal(payload)
	if err != nil {
		app.logger.Error(err.Error(), "location", "helpers.go/json.Marshal")
		return
	}

	w.WriteHeader(status)
	w.Write(body)
}

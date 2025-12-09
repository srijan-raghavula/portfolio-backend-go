package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.sendMessage(&Payload{
		Message: "sweet home alabama",
	}, w, http.StatusOK)
}

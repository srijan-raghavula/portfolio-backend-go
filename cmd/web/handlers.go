package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.sendMessage(&Payload{
		Message: "sweet home alabama",
	}, w, http.StatusOK)
}

func (app *application) tbd(w http.ResponseWriter, r *http.Request) {
	app.sendMessage(&Payload{
		Message: "You ain't dun doin' dis yet.",
	}, w, http.StatusOK)
}

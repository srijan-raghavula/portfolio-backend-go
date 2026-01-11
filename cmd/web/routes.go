package main

import (
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("GET /{$}", app.home)

	// The following endpoints are separated by category
	mux.HandleFunc("POST /v0/sessions", app.tbd)
	mux.HandleFunc("DELETE /v0/sessions", app.tbd)

	mux.HandleFunc("GET /v0/profile", app.tbd)
	mux.HandleFunc("PUT /v0/profile", app.tbd)

	mux.HandleFunc("GET /v0/blogs", app.tbd)
	mux.HandleFunc("GET /v0/blogs/{id}", app.tbd)
	mux.HandleFunc("POST /v0/blogs", app.tbd)
	mux.HandleFunc("PUT /v0/blogs/{id}", app.tbd)
	mux.HandleFunc("DELETE /v0/blogs/{id}", app.tbd)

	mux.HandleFunc("GET /v0/projects", app.tbd)
	mux.HandleFunc("GET /v0/projects/{id}", app.tbd)
	mux.HandleFunc("POST /v0/projects", app.tbd)
	mux.HandleFunc("PUT /v0/projects/{id}", app.tbd)
	mux.HandleFunc("DELETE /v0/projects/{id}", app.tbd)

	std := alice.New(app.recoverPanic, app.logRequest, commonHeaders)
	return std.Then(mux)
}

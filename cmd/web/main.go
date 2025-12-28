package main

import (
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := ":4040"
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app := &application{
		logger: logger,
	}

	err := http.ListenAndServe(addr, app.routes())
	if err != nil {
		app.logger.Error(err.Error())
		os.Exit(1)
	}
}

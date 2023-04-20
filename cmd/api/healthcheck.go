package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	status := envelope{
		"environment": app.config.env,
		"status":      "healthy",
		"port":        fmt.Sprintf("%d", app.config.port),
	}
	err := app.writeJSON(w, http.StatusOK, status, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

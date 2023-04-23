package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := chi.NewRouter()
	router.Get("/v1/healthcheck", app.healthCheckHandler)
	router.Get("/v1/movies/{id}", app.getMovieHandler)
	router.Get("/v1/movies", app.getAllMoviesHandler)
	router.Post("/v1/movies", app.createMovieHandler)
	router.Patch("/v1/movies/{id}", app.updateMovieHandler)
	router.Delete("/v1/movies/{id}", app.deleteMovieHandler)
	return router
}

package main

import (
	"errors"
	"fmt"
	"github.com/rrebeiz/quickmovies/internal/data"
	"github.com/rrebeiz/quickmovies/internal/validator"
	"net/http"
)

func (app *application) getMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidParamID):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	movie, err := app.models.Movies.GetMovie(r.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrNoRecordFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		Runtime int32    `json:"runtime"`
		Year    int32    `json:"year"`
		Genres  []string `json:"genres"`
	}

	err := app.readJSON(w, r, &input)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	movie := &data.Movie{
		Title:   input.Title,
		Runtime: input.Runtime,
		Year:    input.Year,
		Genres:  input.Genres,
	}
	v := validator.NewValidator()

	permittedGenres := []string{"action", "adventure", "comedy", "horror", "drama"}
	data.ValidateMovie(v, movie, permittedGenres...)

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Movies.CreateMovie(r.Context(), movie)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	headers := make(http.Header)

	location := fmt.Sprintf("/v1/movies/%d", movie.ID)

	headers.Set("Location", location)
	err = app.writeJSON(w, http.StatusCreated, envelope{"movie": movie}, headers)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}

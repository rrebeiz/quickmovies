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

func (app *application) getAllMoviesHandler(w http.ResponseWriter, r *http.Request) {
	movies, err := app.models.Movies.GetAllMovies(r.Context())
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"movies": movies}, nil)
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

	data.ValidateMovie(v, movie)

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

func (app *application) updateMovieHandler(w http.ResponseWriter, r *http.Request) {
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

	var input struct {
		Title   *string  `json:"title"`
		Runtime *int32   `json:"runtime"`
		Year    *int32   `json:"year"`
		Genres  []string `json:"genres"`
	}

	err = app.readJSON(w, r, &input)

	if err != nil {
		app.badRequestResponse(w, r, err)
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
	v := validator.NewValidator()

	if input.Title != nil {
		movie.Title = *input.Title
	}

	if input.Year != nil {
		v.Check(*input.Year > 0, "year", "should be a positive number")
		movie.Year = *input.Year
	}

	if input.Runtime != nil {
		v.Check(*input.Runtime > 0, "runtime", "should be a positive number")
		movie.Runtime = *input.Runtime
	}

	if input.Genres != nil {
		v.Check(validator.Unique(input.Genres), "genres", "must not contain duplicate genres")
		permittedGenres := []string{"action", "adventure", "comedy", "horror", "drama"}
		for _, genre := range input.Genres {
			v.Check(validator.PermittedValue(genre, permittedGenres...), "genre", fmt.Sprintf("please use the following genres %s", permittedGenres))
			movie.Genres = input.Genres
		}
	}

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Movies.UpdateMovie(r.Context(), movie)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
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

func (app *application) deleteMovieHandler(w http.ResponseWriter, r *http.Request) {
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
	err = app.models.Movies.DeleteMovie(r.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrNoRecordFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	message := fmt.Sprintf("movie with the id %d has been deleted", id)
	err = app.writeJSON(w, http.StatusOK, envelope{"message": message}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

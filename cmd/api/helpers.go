package main

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type envelope map[string]any

var (
	ErrInvalidParamID = errors.New("invalid param ID")
)

func (app *application) readIDParam(r *http.Request) (int64, error) {
	id, err := strconv.ParseInt(chi.URLParamFromCtx(r.Context(), "id"), 10, 64)
	if err != nil || id <= 0 {
		return 0, ErrInvalidParamID
	}
	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	var js []byte

	if app.config.env == "develop" {
		jsData, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			return err
		}
		js = jsData
	} else {
		jsData, err := json.Marshal(data)
		if err != nil {
			return err
		}
		js = jsData
	}
	js = append(js, '\n')

	for k, v := range headers {
		w.Header()[k] = v
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

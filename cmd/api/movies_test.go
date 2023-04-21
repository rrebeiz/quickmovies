package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetMovieHandler(t *testing.T) {
	tests := []struct {
		name             string
		id               string
		expectedStatus   int
		expectedResponse string
	}{
		{"valid test", "1", http.StatusOK, "{\"movie\":{\"id\":1,\"title\":\"test\",\"runtime\":100,\"year\":2020,\"genres\":[\"action\",\"adventure\"]}}\n"},
		{"not found test", "0", http.StatusNotFound, "{\"error\":\"the requested resource could not be found\"}\n"},
		{"no id test", "", http.StatusNotFound, "{\"error\":\"the requested resource could not be found\"}\n"},
	}

	for _, e := range tests {
		req, _ := http.NewRequest("GET", "/v1/movies/1", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", e.id)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(testApp.getMovieHandler)
		handler.ServeHTTP(rr, req)

		if e.expectedStatus != rr.Code {
			t.Errorf("%s: expected %d but got %d", e.name, e.expectedStatus, rr.Code)
		}

		if e.expectedResponse != rr.Body.String() {
			t.Errorf("%s: expected %s but got %s", e.name, e.expectedResponse, rr.Body.String())
		}
	}
}

func TestCreateMovieHandler(t *testing.T) {
	tests := []struct {
		name             string
		body             string
		expectedStatus   int
		expectedResponse string
	}{
		{"valid test", `{"title":"test","runtime":100,"year":2020,"genres":["action","adventure"]}`, http.StatusCreated, "{\"movie\":{\"id\":2,\"title\":\"test\",\"runtime\":100,\"year\":2020,\"genres\":[\"action\",\"adventure\"]}}\n"},
		{"invalid empty body test", ``, http.StatusBadRequest, "{\"error\":\"body must not be empty\"}\n"},
		{"invalid empty data test", `{"title":"", "runtime":0, "year":0, "genres":[]}`, http.StatusUnprocessableEntity, "{\"error\":{\"genres\":\"should contain at least 1 genre\",\"runtime\":\"should not be empty\",\"title\":\"should not be empty\",\"year\":\"should not be empty\"}}\n"},
	}

	for _, e := range tests {
		req, _ := http.NewRequest("POST", "/v1/movies", strings.NewReader(e.body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(testApp.createMovieHandler)

		handler.ServeHTTP(rr, req)

		if e.expectedStatus != rr.Code {
			t.Errorf("%s: expected %d but got %d", e.name, e.expectedStatus, rr.Code)
		}

		if e.expectedResponse != rr.Body.String() {
			t.Errorf("%s: expected %s but got %s", e.name, e.expectedResponse, rr.Body.String())
		}
	}
}

func TestUpdateMovieHandler(t *testing.T) {
	tests := []struct {
		name             string
		id               string
		body             string
		expectedStatus   int
		expectedResponse string
	}{
		{"valid test", "1", `{"title": "new test","runtime":150,"year":2021,"genres":["action"]}`, http.StatusOK, "{\"movie\":{\"id\":1,\"title\":\"new test\",\"runtime\":150,\"year\":2021,\"genres\":[\"action\"]}}\n"},
		{"not found test", "0", `{"title": "new test","runtime":150,"year":2021,"genres":["action"]}`, http.StatusNotFound, "{\"error\":\"the requested resource could not be found\"}\n"},
		{"validation failed test", "1", `{"runtime":-15,"year":0,"genres":["banana", "banana"]}`, http.StatusUnprocessableEntity, "{\"error\":{\"genre\":\"please use the following genres [action adventure comedy horror drama]\",\"genres\":\"must not contain duplicate genres\",\"runtime\":\"should be a positive number\",\"year\":\"should be a positive number\"}}\n"},
		{"server error test", "2", `{"title": "new test","runtime":150,"year":2021,"genres":["action"]}`, http.StatusInternalServerError, "{\"error\":\"the server encountered a problem and could not process your request\"}\n"},
	}
	for _, e := range tests {
		req, _ := http.NewRequest("PATCH", "/v1/movies/1", strings.NewReader(e.body))
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", e.id)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(testApp.updateMovieHandler)
		handler.ServeHTTP(rr, req)

		if e.expectedStatus != rr.Code {
			t.Errorf("%s: expected %d but got %d", e.name, e.expectedStatus, rr.Code)
		}

		if e.expectedResponse != rr.Body.String() {
			t.Errorf("%s: expected %s but got %s", e.name, e.expectedResponse, rr.Body.String())
		}

	}
}

package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/http/httptest"
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

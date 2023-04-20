package main

import (
	"github.com/rrebeiz/quickmovies/internal/data"
	"os"
	"testing"
)

var testApp application
var testConfig config

func TestMain(m *testing.M) {
	testConfig.env = "production"
	testConfig.port = 4000
	testApp.config = testConfig
	testApp.models = newTestModels()

	os.Exit(m.Run())

}

func newTestModels() data.Models {
	return data.Models{
		Movies: data.NewMockMovieModel(),
	}
}

package data

import (
	"context"
	"errors"
)

type MockMovieModel struct {
}

func NewMockMovieModel() MockMovieModel {
	return MockMovieModel{}
}
func (m MockMovieModel) GetMovie(ctx context.Context, id int64) (*Movie, error) {
	if id == 1 {
		return &Movie{
			ID:      1,
			Title:   "test",
			Runtime: 100,
			Year:    2020,
			Genres:  []string{"action", "adventure"},
			Version: 1,
		}, nil
	} else if id == 0 {
		return nil, ErrNoRecordFound
	} else {
		return nil, errors.New("something went wrong")
	}
}

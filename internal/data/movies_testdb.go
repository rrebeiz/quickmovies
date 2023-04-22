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
		return nil, errors.New("failed to get movie")
	}
}

func (m MockMovieModel) CreateMovie(ctx context.Context, movie *Movie) error {
	if movie.Title == "test" {
		movie.ID = 2
		return nil
	}
	return errors.New("failed to create movie")
}

func (m MockMovieModel) UpdateMovie(ctx context.Context, movie *Movie) error {
	if movie.ID == 1 {
		return nil
	} else if movie.ID == 0 {
		return ErrNoRecordFound
	} else {
		return errors.New("failed to update movie")
	}
}

func (m MockMovieModel) DeleteMovie(ctx context.Context, id int64) error {
	if id == 0 {
		return ErrNoRecordFound
	} else if id == 1 {
		return nil
	} else {
		return errors.New("failed to delete movie")
	}

}

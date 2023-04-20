package data

import (
	"context"
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"time"
)

var (
	ErrNoRecordFound = errors.New("the requested resource could not be found")
)

type Movies interface {
	GetMovie(ctx context.Context, id int64) (*Movie, error)
}

type Movie struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Runtime   int32     `json:"runtime"`
	Year      int32     `json:"year"`
	Genres    []string  `json:"genres"`
	Version   int32     `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type MovieModel struct {
	DB *sql.DB
}

func NewMovieModel(db *sql.DB) MovieModel {
	return MovieModel{DB: db}
}

func (m MovieModel) GetMovie(ctx context.Context, id int64) (*Movie, error) {
	query := `select id, title, runtime, year, genres from movies where id = $1`
	var movie Movie
	if id <= 0 {
		return nil, ErrNoRecordFound
	}
	err := m.DB.QueryRowContext(ctx, query, id).Scan(&movie.ID, &movie.Title, &movie.Runtime, &movie.Year, pq.Array(&movie.Genres))

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNoRecordFound
		default:
			return nil, err
		}
	}
	return &movie, nil
}

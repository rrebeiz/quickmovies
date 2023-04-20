package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"github.com/rrebeiz/quickmovies/internal/validator"
	"time"
)

var (
	ErrNoRecordFound = errors.New("the requested resource could not be found")
)

type Movies interface {
	GetMovie(ctx context.Context, id int64) (*Movie, error)
	CreateMovie(ctx context.Context, movie *Movie) error
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

func (m MovieModel) CreateMovie(ctx context.Context, movie *Movie) error {
	query := `insert into movies (title, runtime, year, genres) values ($1, $2, $3, $4) returning id`
	args := []interface{}{movie.Title, movie.Runtime, movie.Year, pq.Array(movie.Genres)}
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&movie.ID)
	if err != nil {
		return err
	}
	return nil

}

func ValidateMovie(v *validator.Validator, movie *Movie, permittedGenres ...string) {
	v.Check(movie.Title != "", "title", "should not be empty")
	v.Check(len(movie.Title) <= 500, "title", "should not be greater than 500 bytes")

	v.Check(movie.Year != 0, "year", "should not be empty")
	v.Check(movie.Year >= 0, "year", "should be a positive number")

	v.Check(movie.Runtime != 0, "runtime", "should not be empty")
	v.Check(movie.Runtime > 0, "runtime", "should be a positive number")

	v.Check(movie.Genres != nil, "genres", "should not be empty")
	v.Check(len(movie.Genres) >= 1, "genres", "should contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "should not contain more than 5 genres")

	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate genres")

	for _, genre := range movie.Genres {
		v.Check(validator.PermittedValue(genre, permittedGenres...), "genre", fmt.Sprintf("please use the following permitted genres %s", permittedGenres))
	}

}

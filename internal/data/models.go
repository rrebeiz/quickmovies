package data

import (
	"database/sql"
	"errors"
)

var (
	ErrNoRecordFound = errors.New("the requested resource could not be found")
	ErrEditConflict  = errors.New("edit conflict")
)

type Models struct {
	Movies Movies
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies: NewMovieModel(db),
	}
}

package data

import "database/sql"

type Models struct {
	Movies Movies
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies: NewMovieModel(db),
	}
}

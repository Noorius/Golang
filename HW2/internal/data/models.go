package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Knives KnifeModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Knives: KnifeModel{DB: db},
	}
}

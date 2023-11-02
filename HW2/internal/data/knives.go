package data

import (
	"database/sql"
	"errors"
	"hw2.nur.net/internal/validator"
	"time"
)

type Knife struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // Always hides
	Title     string    `json:"title"`
	Material  string    `json:"material"`
	Color     string    `json:"color"`
	Country   string    `json:"country,omitempty"` // Hides if empty
	Duration  Duration  `json:"duration,omitempty"`
	Version   int       `json:"version"`
}

func ValidateKnife(v *validator.Validator, knife *Knife) {
	v.Check(knife.Title != "", "title", "must be provided")
	v.Check(len(knife.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(knife.Material != "", "material", "must be provided")
	v.Check(knife.Color != "", "color", "must be provided")
	v.Check(knife.Country != "", "country", "must be provided")
	v.Check(knife.Duration != 0, "duration", "must be provided")
	v.Check(knife.Duration > 0, "duration", "must be a positive integer")
}

type KnifeModel struct {
	DB *sql.DB
}

func (k KnifeModel) Insert(kn *Knife) error {
	query := `INSERT INTO knives (title, material, color, country, duration) 
				VALUES ($1, $2, $3, $4, $5) 
				RETURNING id, created_at`

	args := []interface{}{kn.Title, kn.Material, kn.Color, kn.Country, kn.Duration}

	return k.DB.QueryRow(query, args...).Scan(&kn.ID, &kn.CreatedAt)
}

func (k KnifeModel) Get(id int64) (*Knife, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, created_at, title, material, color, country, duration, version
		FROM knives
		WHERE id = $1`

	var knife Knife

	err := k.DB.QueryRow(query, id).Scan(
		&knife.ID,
		&knife.CreatedAt,
		&knife.Title,
		&knife.Material,
		&knife.Color,
		&knife.Country,
		&knife.Duration,
		&knife.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &knife, nil
}

func (k KnifeModel) Update(kn *Knife) error {
	query := `
		UPDATE knives
		SET title = $1, material = $2, color = $3, country = $4, duration = $5, version = version + 1
		WHERE id = $6 AND version = $7
		RETURNING version`

	args := []interface{}{
		kn.Title,
		kn.Material,
		kn.Color,
		kn.Country,
		kn.Duration,
		kn.ID,
		kn.Version,
	}

	err := k.DB.QueryRow(query, args...).Scan(&kn.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}

func (k KnifeModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM knives
		WHERE id = $1`

	result, err := k.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}

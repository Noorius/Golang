package data

import (
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

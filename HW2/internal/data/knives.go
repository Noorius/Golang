package data

import "time"

type Knife struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // Always hides
	Title     string    `json:"title"`
	Material  string    `json:"material"`
	Color     string    `json:"color"`
	Country   string    `json:"country,omitempty"` // Hides if empty
}

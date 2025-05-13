package entity

import "time"

type Product struct {
	ID              int
	Name            string
	Description     string
	Price           float64
	ImageURL        string
	PreparationTime int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CategoryID      int
}

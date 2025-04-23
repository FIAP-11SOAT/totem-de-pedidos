package entity

import "time"

type CategoryProduct struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
}

type Product struct {
	ID              int
	Name            string
	Description     string
	Price           float64
	ImageURL        string
	PreparationTime int
	CreatedAt       time.Time
	CategoryID      int
	Category        *CategoryProduct
}

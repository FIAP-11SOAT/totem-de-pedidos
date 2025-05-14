package entity

import "time"

type ProductCategory struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Product struct {
	ID              int     `json:"id"`
	Name            string  `json:"nome"`
	Description     string  `json:"descricao"`
	Price           float64 `json:"preco"`
	ImageURL        string  `json:"imagem"`
	PreparationTime int     `json:"preparation_time"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CategoryID      int `json:"categoriaId"`
}

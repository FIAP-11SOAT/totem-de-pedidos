package dto

type ProductResponse struct {
	ID              string  `json:"id"`
	Name            string  `json:"nome"`
	Description     string  `json:"descricao"`
	Image           string  `json:"imagem"`
	Price           float64 `json:"preco"`
	PreparationTime int     `json:"preparation_time"`
	CategoryID      int     `json:"categoriaId"`
}

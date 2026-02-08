package model

import "github.com/google/uuid"

type Product struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Price    float64   `json:"price"`
	Quantity int       `json:"quantity"`
	Category Category  `json:"category"`
	Status   bool      `json:"status"`
}

type ProductDto struct {
	Name       string    `json:"name"`
	Price      float64   `json:"price"`
	Quantity   int       `json:"quantity"`
	CategoryID uuid.UUID `json:"category_id"`
	Status     bool      `json:"status"`
}

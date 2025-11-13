package models

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price" binding:"required"`
	Stock       int     `json:"stock"`
}

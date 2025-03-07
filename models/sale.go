package models

import "gorm.io/gorm"

type Sale struct {
	gorm.Model
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
}

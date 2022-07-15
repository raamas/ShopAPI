package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	ProductID       uint      `gorm:"unique_index;not_null" json:"product_id"`
	ProductName     string    `gorm:"not_null;" json:"product_name"`
	ProductCode     uuid.UUID `gorm:"not_null;default:uuid_generate_v4()" json:"product_code"`
	ProductCategory string    `json:"product_category"`
	ProductPrice    uint      `gorm:"not_null" json:"product_price"`
}

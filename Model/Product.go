package Model

import (
	"time"

	_ "gorm.io/gorm"
)

type Product struct {
    ID        uint      `gorm:"primarykey" json:"id"`
    Name      string    `form:"name" json:"name" binding:"required"`
    Price     float64   `form:"price" json:"price" binding:"required"`
    Stock     int       `form:"stock" json:"stock" binding:"required"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
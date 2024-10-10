package Model

import (
	"time"

	_ "gorm.io/gorm"
)

type Operator struct {
    ID        uint      `gorm:"primarykey" json:"id"`
    Username  string    `form:"username" json:"username" binding:"required"`
    Password  string    `form:"password" json:"password" binding:"required"`
    CreatedAt time.Time `json:"created_at"`
}
package Model

import (
	"time"
)

type Operator struct {
    ID        uint      `gorm:"primarykey" json:"id"`
    Name  string    `gorm:"unique;not null" form:"name" json:"name" binding:"required"` 
    Password  string    `form:"password" json:"password" binding:"required"`
    CreatedAt time.Time `json:"created_at"`
}
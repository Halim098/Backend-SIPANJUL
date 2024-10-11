package Model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Operator struct {
    ID        uint      `gorm:"primarykey" json:"id"`
    Name  string    `gorm:"unique;not null" form:"name" json:"name" binding:"required"` 
    Password  string    `form:"password" json:"password" binding:"required"`
    CreatedAt time.Time `json:"created_at"`
}

type OperatorLogin struct {
	Name string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (o *Operator) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(o.Password), []byte(password))
}


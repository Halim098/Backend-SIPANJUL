package Model

import (
	"Sipanjul/Database"
	"errors"
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

func FindOperator (name, password string) (string,error) {
	var Operator Operator
	err := Database.Database.Raw("SELECT * FROM operators WHERE name = ? AND password = ?", name, password).Scan(&Operator)
	if err.Error != nil {
		return "", err.Error
	}

	if err.RowsAffected == 0 {
		return "", errors.New("user not found")
	}

	return Operator.Name, nil
}

func GetOperator () ([]Operator,error) {
    var Operator []Operator
    
    err := Database.Database.Raw("SELECT * FROM operators)").Scan(&Operator)
    if err.Error != nil {
        return Operator, err.Error
    }

    if err.RowsAffected == 0 {
        return Operator, errors.New("user not found")
    }

    return Operator, nil
}
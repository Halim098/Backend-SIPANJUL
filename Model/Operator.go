package Model

import (
	"Sipanjul/Database"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Operator struct {
    ID        uint      `gorm:"primarykey" json:"id"`
    Name  string    `gorm:"unique;not null" form:"name" json:"name"` 
    Password  string    `form:"password" json:"password"`
    Status    bool    `json:"status"`
    CreatedAt time.Time `json:"created_at"`
}

type OperatorLogin struct {
	Name string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type StatusStore struct {
    Storestatus bool `json:"storestatus,omitempty"`
}

type NameOperator struct {
    Name string `json:"name"`
}

func (o *Operator) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(o.Password), []byte(password))
}

func (o *Operator) BeforeSave(*gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(o.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	o.Password = string(hashedPassword)
	return nil
}

func (o *Operator) Save() error {
	err := o.BeforeSave(Database.Database)
	if err != nil {
		return err
	}

	result, _ := FindOperatorByName(o.Name)
	if result.Name == o.Name {
		return errors.New("username already exists")
	}

	o.CreatedAt = time.Now()
    o.Status = true

	err = Database.Database.Exec("INSERT INTO operators (name, password, created_at) VALUES (?, ?, ?)", o.Name, o.Password, o.CreatedAt).Error
	if err != nil {
		return errors.New("failed to create user")
	}
	
	return nil
}

func FindOperatorByName(username string) (Operator, error) {
	var opr Operator
	err := Database.Database.Raw("SELECT * FROM operators WHERE name = ?", username).Scan(&opr)
	if err.Error != nil {
		return Operator{}, err.Error
	}

	if err.RowsAffected == 0 {
		return Operator{}, errors.New("user not found")
	}

	return opr, nil
}

func FindOperatorByID(id uint) (NameOperator, error) {
	var opr NameOperator
	err := Database.Database.Raw("SELECT name FROM operators WHERE id = ?", id).Scan(&opr)
	if err.Error != nil {
		return NameOperator{}, err.Error
	}

	if err.RowsAffected == 0 {
		return NameOperator{}, errors.New("user not found")
	}

	return opr, nil
}

func UpdateOperatorStatus(id uint, status bool) error {
    err := Database.Database.Exec("UPDATE operators SET status = ? WHERE id = ?", status, id).Error
    if err != nil {
        return errors.New("failed to update status")
    }
    return nil
}

func GetStatus(id uint) (StatusStore, error) {
    var status StatusStore
    err := Database.Database.Raw("SELECT status as storestatus FROM operators WHERE id = ?", id).Scan(&status)
    if err.Error != nil {
        return StatusStore{}, err.Error
    }

    if err.RowsAffected == 0 {
        return StatusStore{}, errors.New("data not found")
    }

    return status, nil
}
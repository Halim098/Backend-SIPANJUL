package Model

import (
	"Sipanjul/Database"
	"time"

	_ "gorm.io/gorm"
)

type Product struct {
    ID        uint      `gorm:"primarykey" json:"id"`
    Name      string    `form:"name" json:"name" binding:"required"`
    Price     float64   `form:"price" json:"price" binding:"required"`
    Stock     int       `form:"stock" json:"stock" binding:"required"`
	Division  string	`form:"division" json:"division" binding:"required"`
	Location  string	`form:"location" json:"location" binding:"required"`
	Active	  string	`form:"active" json:"active" binding:"required"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
	DeleteAt  time.Time `json:"delete_at"`
}

func AddProduct (name string, price float64, stock int, division string, location string) error {
	Product := Product{Name: name, Price: price, Stock: stock, Division: division, Location: location, Active: "true", CreatedAt: time.Now()}
	err := Database.Database.Create(&Product)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func UpdateProduct (id uint, name string, price float64, stock int, division string, location string) error {
	Product := Product{ID: id}
	err := Database.Database.First(&Product)
	if err.Error != nil {
		return err.Error
	}
	Product.Name = name
	Product.Price = price
	Product.Stock = stock
	Product.Division = division
	Product.Location = location
	Product.UpdatedAt = time.Now()
	err = Database.Database.Save(&Product)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func DeleteProduct (id uint) error {
	Product := Product{ID: id}
	err := Database.Database.First(&Product)
	if err.Error != nil {
		return err.Error
	}
	Product.Active = "false"
	Product.DeleteAt = time.Now()
	err = Database.Database.Save(&Product)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func GetProduct () ([]Product,error) {
	var Product []Product
	
	err := Database.Database.Raw("SELECT * FROM products WHERE active = ?", "true").Scan(&Product)
	if err.Error != nil {
		return Product, err.Error
	}

	if err.RowsAffected == 0 {
		return Product, nil
	}

	return Product, nil
}
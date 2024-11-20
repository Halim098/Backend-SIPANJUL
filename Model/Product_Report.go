package Model

import (
	"Sipanjul/Database"
	"time"

	_ "gorm.io/gorm"
)

var query string = "SELECT r.id, p.name, r.quantity, r.action, r.description r.date FROM product_reports r JOIN products p ON r.prod_id = p.id"

type ProductReport struct {
	ID        	uint      	`gorm:"primarykey" json:"id"`
	ProdID	    uint      	`json:"prod_id"`
	Quantity  	int       	`json:"quantity" binding:"required"`
	Action		string    	`json:"action" binding:"required"`
	Description	string		`json:"description" binding:"required"`
	Date 		time.Time 	`json:"date"`

	Product     Product     `gorm:"foreignKey:ProdID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Foreign key defined
}

func AddProductReport (data *ProductReport) error {
	err := Database.Database.Create(&data)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
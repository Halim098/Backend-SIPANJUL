package Model

import (
	"Sipanjul/Database"
	"errors"
	"time"

	_ "gorm.io/gorm"
)

type Product_repost struct {
	ID        	uint      	`gorm:"primarykey" json:"id"`
	ProdID	    uint      	`json:"prod_id"`
	Quantity  	int       	`json:"quantity" binding:"required"`
	Action		string    	`json:"action" binding:"required"`

	Product     Product     `gorm:"foreignKey:ProdID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Foreign key defined
}


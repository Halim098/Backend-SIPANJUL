package Model

import (
	"Sipanjul/Database"

	_ "gorm.io/gorm"
)

type Sales_Detail struct {
    ID        uint      `gorm:"primarykey" json:"id"`
    SalesID   uint      `json:"sales_id"`
    ProdID    uint      `json:"prod_id" binding:"required"`
    Quantity  int       `json:"quantity" binding:"required"`
    Total     int       `json:"total" binding:"required"`
    
    Product   Product   `gorm:"foreignKey:ProdID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Foreign key defined
}

func AddSalesDetail (data *Sales_Detail) error {
	err := Database.Database.Create(&data)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
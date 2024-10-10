package Model

import (
	_ "gorm.io/gorm"
)

type Sales_Detail struct {
    ID        uint      `gorm:"primarykey" json:"id"`
    SalesID   uint      `json:"sales_id"` // Foreign key for Sales
    ProdID    uint      `json:"prod_id"`  // Foreign key for Product
    Quantity  int       `json:"quantity"`
    Total     int       `json:"total"`
    
    Product   Product   `gorm:"foreignKey:ProdID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Foreign key defined
}
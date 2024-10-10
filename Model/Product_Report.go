package Model

import (
	"time"

	_ "gorm.io/gorm"
)

type Product_Report struct {
    ID        uint      `gorm:"primarykey" json:"id"`
    ProdID    uint      `json:"prod_id"`  // Foreign key for Product
    OprID     uint      `json:"opr_id"`   // Foreign key for Operator
    Action    string    `json:"action"`
    Quality   int       `json:"quality"`
    Note      string    `json:"note"`
    Date      time.Time `json:"date"`
    
    Operator  Operator  `gorm:"foreignKey:OprID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Foreign key defined
    Product   Product   `gorm:"foreignKey:ProdID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Foreign key defined
}
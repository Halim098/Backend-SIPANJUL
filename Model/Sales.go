package Model

import (
	"Sipanjul/Database"
	"time"

	_ "gorm.io/gorm"
)

type Sales struct {
    ID        uint      `gorm:"primarykey" json:"id"`
    OprID     uint      `json:"opr_id"`  // Foreign key for Operator
    Total     int       `json:"total"`
    Date      time.Time `json:"date"`
    
    Operator  Operator  `gorm:"foreignKey:OprID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Foreign key defined
}

func AddSales (data *Sales) error {
	err := Database.Database.Create(&data)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
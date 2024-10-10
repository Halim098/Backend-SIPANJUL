package Model

import (
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
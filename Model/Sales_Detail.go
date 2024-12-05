package Model

import (
	"Sipanjul/Database"
	"errors"

	_ "gorm.io/gorm"
)

type Sales_Detail struct {
    ID        uint      `gorm:"primarykey" json:"id"`
    SalesID   uint      `json:"sales_id"`
    ProdID    uint      `json:"prod_id" binding:"required"`
    Quantity  int       `json:"quantity" binding:"required"`
	StockAwal int       `json:"stockawal" binding:"required"`
	StockAkhir int      `json:"stockakhir" binding:"required"`
    Total     int       `json:"total" binding:"required"`
    
    Product   Product   `gorm:"foreignKey:ProdID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Foreign key defined
	Sales     Sales     `gorm:"foreignKey:SalesID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Foreign key defined
}

func AddSalesDetail (data *Sales_Detail) error {
	err := Database.Database.Create(&data)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func GetSalesDetailbySalesandDate (idsales uint,sartdate,endate string) ([]Sales_Detail, error) {
	var Sales []Sales_Detail

	err := Database.Database.Raw(`SELECT 
            sd.*, 
            s.date AS sales_date, 
            p.division AS product_division, 
            p.name AS product_name
        FROM 
            sales_details sd
        JOIN 
            sales s ON sd.sales_id = s.id
        JOIN 
            products p ON sd.prod_id = p.id
        WHERE 
            s.id = ? AND 
            s.date BETWEEN ? AND ? AND
			p.active = ?`, idsales, sartdate, endate,"true").Scan(&Sales)
	if err.Error != nil {
		return Sales, err.Error
	}

	if err.RowsAffected == 0 {
		return Sales, errors.New("data tidak ditemukan")
	}

	return Sales, nil
}


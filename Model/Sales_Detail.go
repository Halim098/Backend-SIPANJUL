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

type BestSelling struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Stock bool `json:"stock"`
	Packagesize string `json:"packagesize"`
	Types string `json:"type"`
	Imageurl string `json:"imageurl"`
}

type LastTransaction struct {
	Name string `json:"name"`
	Kuantitas int `json:"kuantitas"`
	Total int `json:"total"`
	Date string `json:"date"`
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
            s.date, 
            p.division, 
            p.name
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

func GetLastTransaction (oprid uint) ([]LastTransaction, error) {
	var LastTrans []LastTransaction

	err := Database.Database.Raw(`SELECT 
			p.name, 
			sd.quantity, 
			sd.total, 
			s.date
		FROM 
			sales_details sd
		JOIN 
			sales s ON sd.sales_id = s.id
		JOIN 
			products p ON sd.prod_id = p.id
		WHERE 
			s.opr_id = ?
		ORDER BY 
			s.date DESC
		LIMIT 5`, oprid).Scan(&LastTrans)
	if err.Error != nil {
		return LastTrans, err.Error
	}

	if err.RowsAffected == 0 {
		return LastTrans, errors.New("data tidak ditemukan")
	}

	return LastTrans, nil
}

func GetBestSellingItem(oprid uint) ([]Sales_Detail, error) {
	var BestSelling []Sales_Detail

	err := Database.Database.Raw(`SELECT 
			p.id,
			p.name,
			p.packagesize,
			p.type,
			p.imageurl,
			p,stock
		FROM 
			sales_details sd
		JOIN 
			products p ON sd.prod_id = p.id
		JOIN 
			sales s ON sd.sales_id = s.id
		WHERE 
			s.opr_id = ?
		GROUP BY 
			p.name
		ORDER BY 
			SUM(sd.quantity) DESC LIMIT 10`, oprid).Scan(&BestSelling)
	if err.Error != nil {
		return BestSelling, err.Error
	}

	if err.RowsAffected == 0 {
		return BestSelling, errors.New("data tidak ditemukan")
	}

	return BestSelling, nil
}

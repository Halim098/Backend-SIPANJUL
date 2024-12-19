package Model

import (
	"Sipanjul/Database"
	"errors"
	"strings"
	"time"

	_ "gorm.io/gorm"
)

type ProductReport struct {
	ID        	uint      	`gorm:"primarykey" json:"id"`
	ProdID	    uint      	`json:"prod_id"`
	Quantity  	int       	`json:"quantity" binding:"required"`
	Action		string    	`json:"action" binding:"required"`
	Description	string		`json:"description" binding:"required"`
	Date 		time.Time 	`json:"date"`

	Product     Product     `gorm:"foreignKey:ProdID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Foreign key defined
}

type ProductReportOpr struct {
    ID        	uint      	`json:"id"`
    Name      	string    	`json:"name" binding:"required"`
    Quantity  	int       	`json:"quantity" binding:"required"`
    Action		string    	`json:"action" binding:"required"`
    Description	string		`json:"description" binding:"required"`
    Date 		time.Time 	`json:"date"`
}

func AddProductReport (data *ProductReport) error {
	err := Database.Database.Create(&data)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func GetProductReport (startdate,enddate,divisi,detail string) ([]ProductReportOpr, error) {
	var Product []ProductReportOpr

	query := gettingQuery(startdate,enddate,divisi,detail)
	
	err := Database.Database.Raw(query).Scan(&Product)
	if err.Error != nil {
		return Product, err.Error
	}

	if err.RowsAffected == 0 {
		return Product, errors.New("data tidak ditemukan")
	}

	return Product, nil
}

func gettingQuery(startdate, enddate, divisi, detail string) string {
	query := `
        SELECT 
            r.id, 
            p.name, 
            r.quantity, 
            r.action, 
            r.description, 
            r.date 
        FROM 
            product_reports r 
        JOIN 
            products p 
        ON 
            r.prod_id = p.id 
    `

	conditions := []string{}

	if startdate != "" {
		if enddate != "" {
			conditions = append(conditions, "r.date >= '"+startdate+"' AND r.date < '"+enddate+"'")
		} else {
			conditions = append(conditions, "r.date = '"+startdate+"'")
		}
	}
	if divisi != "" {
		conditions = append(conditions, "p.division = '"+divisi+"'")
	}
	if detail != "" {
		conditions = append(conditions, "r.action = '"+detail+"'")
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	return query
}

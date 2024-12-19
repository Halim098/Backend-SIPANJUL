package Model

import (
	"Sipanjul/Database"
	"errors"
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
	var Product []ProductReport
    var ProductOpr []ProductReportOpr

	query := gettingQuery(startdate,enddate,divisi,detail)
	
	err := Database.Database.Raw(query).Scan(&Product)
	if err.Error != nil {
		return ProductOpr, err.Error
	}

	if err.RowsAffected == 0 {
		return ProductOpr, errors.New("data tidak ditemukan")
	}

    for _, data := range Product {
        ProductOpr = append(ProductOpr, ProductReportOpr{
            ID: data.ID,
            Name: data.Product.Name,
            Quantity: data.Quantity,
            Action: data.Action,
            Description: data.Description,
            Date: data.Date,
        })
    }

	return ProductOpr, nil
}

func gettingQuery (startdate,enddate,divisi,detail string) string {
	if startdate == "" && divisi == "" && detail == "" {
		return query
	}
	query = query + " WHERE "
	if startdate != "" {
		if enddate != "" {
			query = query + "r.date >= " + startdate + " AND r.date <=" + enddate + " AND "
		} else {
			query = query + "r.date = " + startdate + " AND "
		}
	}
	if divisi != "" {
		query = query + "p.division = " + divisi + " AND "
	}
	if detail != "" {
		query = query + "r.action = " + detail + " AND "
	}
	return query[:len(query)-5]
}
package Model

import (
	"time"

	_ "gorm.io/gorm"
)

type Sales_Report struct {
	ID        	uint      	`gorm:"primarykey" json:"id"`
	SalesID	    uint      	`json:"sales_id"`
	Stockawal  	int       	`json:"stockawal" binding:"required"`
	Komoditi   	string    	`json:"komoditi" binding:"required"`
	Terjual  	int       	`json:"terjual" binding:"required"`
	Price		int    		`json:"hargapenjualan" binding:"required"`
	Stockakhir	int       	`json:"stockakhir" binding:"required"`
}

func GetSalesDetail(id uint, startdate, endate string) ([]Sales_Report, error) {
    var Sales []Sales_Report
    var saleMap = make(map[string]Sales_Report)

    start,err := time.Parse("2006-01-02", startdate)
    if err != nil {
        return Sales, err
    }

    end,err := time.Parse("2006-01-02", endate)
    if err != nil {
        return Sales, err
    }
    
    data, err := GetSalesDetailbySalesandDate(id, startdate, endate)
    if err != nil {
        return Sales, err
    }

    for _, v := range data {
        if existing, found := saleMap[v.Product.Name]; found {
			// Jika produk dengan nama ini sudah ada, tambahkan kuantitas
			existing.Price += v.Total
			saleMap[v.Product.Name] = existing

            if v.Sales.Date.Equal(start) {
                if v.StockAwal > existing.Stockawal {
                    existing.Stockawal = v.StockAwal
                }
            }

            if v.Sales.Date.Equal(end) {
                if v.StockAkhir < existing.Stockakhir {
                    existing.Stockakhir = v.StockAkhir
                }
            }

		} else {
            saleMap[v.Product.Name] = Sales_Report{
                ID: v.ID,
                SalesID: v.SalesID,
                Stockawal: v.StockAwal,
                Komoditi: v.Product.Name,
                Terjual: v.Quantity,
                Price: v.Total,
                Stockakhir: v.StockAkhir,
            }
        }
    }

    for _, v := range saleMap {
        Sales = append(Sales, v)
    }

    return Sales, nil
}
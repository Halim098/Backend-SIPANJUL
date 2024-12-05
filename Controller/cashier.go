package Controller

import (
	"Sipanjul/Model"
)

func Checkout(oprid uint, total int, detail *[]Model.Sales_Detail) error {
	data := Model.Sales{OprID: oprid, Total: total}
	err := Model.AddSales(&data)
	if err != nil {
		return err
	}

	for _, v := range *detail {
		v.SalesID = data.ID
		prodData, err := GetProductByID(v.ProdID)
		if err != nil {
			return err
		}
		v.StockAwal = prodData.Stock
		v.StockAkhir = v.StockAwal - v.Quantity

		err = Model.AddSalesDetail(&v)
		if err != nil {
			return err
		}

		updated := Model.Product{
			Division: prodData.Division,
			Name: prodData.Name,
			Price: prodData.Price,
			Stock: -v.Quantity,
			UpdatedAt: prodData.UpdatedAt,
		}

		err = Model.UpdateProduct(&prodData, &updated)
		if err != nil {
			return err
		}
	
	}
	return nil
}
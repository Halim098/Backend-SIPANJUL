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

		err = Model.UpdateStock(v.ProdID, v.StockAkhir)
		if err != nil {
			return err
		}
	}
	return nil
}
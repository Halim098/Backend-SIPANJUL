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
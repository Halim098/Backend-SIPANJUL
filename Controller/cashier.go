package Controller

import (
	"Sipanjul/Model"
	"errors"
)

func Checkout(oprid uint, total int, detail *[]Model.Sales_Detail) error {
	data := Model.Sales{OprID: oprid, Total: total}
	err := Model.AddSales(&data)
	if err != nil {
		return err
	}

	id,err := Model.Get1lasttransaction()
	if err != nil {
		return err
	}
	if id == 0 {
		return errors.New("data tidak ditemukan")
	}

	for _, v := range *detail {
		v.SalesID = id
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
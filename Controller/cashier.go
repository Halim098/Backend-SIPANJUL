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
		err = Model.AddSalesDetail(&v)
		if err != nil {
			return err
		}
	}
	return nil
}
package Controller

import (
	"Sipanjul/Model"
	"sync"
)

func GetLastTransaction(oprid uint) ([]Model.LastTransaction, error) {
	data, err := Model.GetLastTransaction(oprid)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetBestSellingItem(oprid uint) ([]Model.BestSelling, []Model.BestSelling,error,error) {
	var weekly []Model.BestSelling
	var Weeklyerr error

	var monthly []Model.BestSelling
	var Monthlyerr error

	var wg1 sync.WaitGroup
	wg1.Add(2)
	go func() {
		defer wg1.Done()
		data, err := Model.GetBestSellingItemWeekly(oprid)
		if err != nil {
			Weeklyerr = err
		}

		for _, v := range data {
			var ready bool
			if v.Stock > 0 {
				ready = true
			}
			if v.Stock <= 0 {
				ready = false
			}
			weekly = append(weekly, Model.BestSelling{
				ID: v.ID,
				Name: v.Name,
				Packagesize: v.Packagesize,
				Types: v.Type,
				Imageurl: v.Imageurl,
				Stock: ready,
			})
		}
	}()

	go func() {
		defer wg1.Done()
		data, err := Model.GetBestSellingItemMonthly(oprid)
		if err != nil {
			Monthlyerr = err
		}

		for _, v := range data {
			var ready bool
			if v.Stock > 0 {
				ready = true
			}
			if v.Stock <= 0 {
				ready = false
			}
			monthly = append(monthly, Model.BestSelling{
				ID: v.ID,
				Name: v.Name,
				Packagesize: v.Packagesize,
				Types: v.Type,
				Imageurl: v.Imageurl,
				Stock: ready,
			})
		}
	}()

	return weekly, monthly, Weeklyerr, Monthlyerr
}
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

func GetBestSellingItem(oprid uint) ([]Model.BestSelling, []Model.BestSelling, error, error) {
    var weekly, monthly []Model.BestSelling
    var Weeklyerr, Monthlyerr error
    var mu sync.Mutex
    var wg1 sync.WaitGroup

    wg1.Add(2)
    go func() {
        defer wg1.Done()
        data, err := Model.GetBestSellingItemWeekly(oprid)
        if err != nil {
            mu.Lock()
            Weeklyerr = err
            mu.Unlock()
            return
        }
        weekly = processBestSelling(data)
    }()

    go func() {
        defer wg1.Done()
        data, err := Model.GetBestSellingItemMonthly(oprid)
        if err != nil {
            mu.Lock()
            Monthlyerr = err
            mu.Unlock()
            return
        }
        monthly = processBestSelling(data)
    }()

    wg1.Wait()
    return weekly, monthly, Weeklyerr, Monthlyerr
}

func processBestSelling(data []Model.BestSellingDB) []Model.BestSelling {
    var result []Model.BestSelling
    for _, v := range data {
		var ready bool
        if v.Stock > 0 {
			ready = true
		} else {
			ready = false
		}

        result = append(result, Model.BestSelling{
            ID:          v.ID,
            Name:        v.Name,
            Packagesize: v.Packagesize,
            Types:       v.Type,
            Imageurl:    v.Imageurl,
            Stock:       ready,
        })
    }
    return result
}

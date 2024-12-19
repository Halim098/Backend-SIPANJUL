package Controller

import "Sipanjul/Model"

func GetLastTransaction(oprid uint) ([]Model.LastTransaction, error) {
	data, err := Model.GetLastTransaction(oprid)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetBestSellingItem(oprid uint) ([]Model.BestSelling, error) {
	var best []Model.BestSelling
	data, err := Model.GetBestSellingItem(oprid)
	if err != nil {
		return nil, err
	}

	for _, v := range data {
		var ready bool
		if v.Stock > 0 {
			ready = true
		}
		if v.Stock <= 0 {
			ready = false
		}
		best = append(best, Model.BestSelling{
			ID: v.ID,
			Name: v.Name,
			Packagesize: v.Packagesize,
			Types: v.Type,
			Imageurl: v.Imageurl,
			Stock: ready,
		})
	}
	return best, nil
}
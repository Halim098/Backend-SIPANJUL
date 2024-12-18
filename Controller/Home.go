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
		if v.Product.Stock > 0 {
			ready = true
		}
		if v.Product.Stock <= 0 {
			ready = false
		}
		best = append(best, Model.BestSelling{
			ID: v.Product.ID,
			Name: v.Product.Name,
			Packagesize: v.Product.Packagesize,
			Types: v.Product.Type,
			Imageurl: v.Product.Imageurl,
			Stock: ready,
		})
	}
	return best, nil
}
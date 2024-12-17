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
	data, err := Model.GetBestSellingItem(oprid)
	if err != nil {
		return nil, err
	}
	return data, nil
}
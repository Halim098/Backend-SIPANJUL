package Controller

import (
	"Sipanjul/Model"
	"time"
)

func AddProduct(data *Model.Product) error {
	err := Model.AddProduct(data)
	if err != nil {
		return err
	}
	return nil
}

func UpdateProduct(stok, data *Model.Product) error {
	data.UpdatedAt = time.Now()
	err := Model.UpdateProduct(stok, data)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProduct(data *Model.Product) error {
	err := Model.DeleteProduct(data)
	if err != nil {
		return err
	}
	return nil
}

func GetProductBYOpr(opr uint) ([]Model.ProductOperator, error) {
	data, err := Model.GetProductByOpr(opr)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetProductByID(id uint) (Model.Product, error){
	data,err := Model.GetProductByID(id)
	if err != nil {
		return Model.Product{}, err
	}

	return data, nil
}

func GetAllProduct() ([]Model.ProductUser, error) {
	prod,err := Model.GetAllProduct()
	if err != nil {
		return []Model.ProductUser{}, err
	}

	newProd := []Model.ProductUser{}
	ready := false

	for _,v := range prod {
		if v.Stock > 0 {
			ready = true
		} else {
			ready = false
		}

		newProd = append(newProd, Model.ProductUser{
			ID: v.ID,
			Name: v.Name,
			Stock: ready,
			Packagesize: v.Packagesize,
			Imageurl: v.Imageurl,
		})
	}

	return newProd, nil
}

func UpdateStock(data *Model.Product, stock int) error {
	data.Stock = data.Stock + stock
	err := Model.UpdateStock(data.ID, data.Stock)
	if err != nil {
		return err
	}
	return nil
}
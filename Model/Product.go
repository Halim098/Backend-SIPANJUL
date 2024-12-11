package Model

import (
	"Sipanjul/Database"
	"errors"
	"time"

	_ "gorm.io/gorm"
)

type Product struct {
    ID        	uint      	`gorm:"primarykey" json:"id"`
    Name      	string    	`json:"name" binding:"required"`
    Price     	int		   	`json:"price" binding:"required"`
    Stock     	int       	`json:"stock" binding:"required"`
	Packagesize	string		`json:"packagesize" binding:"required"`
	Division  	string		`json:"division" binding:"required"`
	Imageurl	string		`json:"image_url" binding:"required"`
	OprID	  	uint	    `json:"opr_id"`
	Active	  	string		`json:"active"`
    CreatedAt 	time.Time 	`json:"created_at"`
    UpdatedAt 	time.Time 	`json:"updated_at"`
	DeleteAt  	time.Time 	`json:"delete_at"`

	Operator  Operator  `gorm:"foreignKey:OprID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type ProductUser struct {
    ID        	uint      	`json:"id"`
    Name      	string    	`json:"name" binding:"required"`
    Stock     	bool       	`json:"stock" binding:"required"`
	Packagesize	string		`json:"packagesize" binding:"required"`
	Imageurl	string		`json:"image_url" binding:"required"`
}

type ProductOperator struct {
    ID        	uint      	`json:"id"`
    Name      	string    	`json:"name" binding:"required"`
    Price     	int		   	`json:"price" binding:"required"`
    Stock     	int       	`json:"stock" binding:"required"`
    Packagesize	string		`json:"packagesize" binding:"required"`
    Division  	string		`json:"division" binding:"required"`
    Imageurl	string		`json:"image_url" binding:"required"`
}

func AddProduct (data *Product) error {
	data.Active = "true"
	err := Database.Database.Create(&data)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func UpdateProduct (stok *Product, data *Product) error {
	stok.Division = data.Division
	stok.Name = data.Name
	stok.Price = data.Price
	stok.Stock = data.Stock
	stok.UpdatedAt = data.UpdatedAt

	err := Database.Database.Save(&stok)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func DeleteProduct (product *Product) error {
	err := Database.Database.First(&product)
	if err.Error != nil {
		return err.Error
	}
	product.Active = "false"
	product.DeleteAt = time.Now()
	err = Database.Database.Save(&product)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func GetProductByOpr (opr uint) ([]ProductOperator,error) {
	var Product []ProductOperator
	
	err := Database.Database.Raw("SELECT id, name, price, stock, packagesize, division, imageurl FROM products WHERE active = ? AND opr_id = ?", "true", opr).Scan(&Product)
	if err.Error != nil {
		return Product, err.Error
	}

	if err.RowsAffected == 0 {
		return Product, errors.New("data tidak ditemukan")
	}

	return Product, nil
}

func GetProductByID (id uint) (Product,error)  {
	var Product Product

	err := Database.Database.Raw("SELECT id, name, price, stock, packagesize, division, imageurl, opr_id FROM products WHERE active = ? AND id = ?", "true", id).Scan(&Product)
	if err.Error != nil {
		return Product, err.Error
	}

	if err.RowsAffected == 0 {
		return Product, errors.New("data tidak ditemukan")
	}

	return Product, nil
}

func GetAllProduct ()([]Product,error){
	var Product []Product
	
	err := Database.Database.Raw("SELECT id, name, stock, packagesize, division, imageurl FROM products WHERE active = ? ", "true").Scan(&Product)
	if err.Error != nil {
		return Product, err.Error
	}

	if err.RowsAffected == 0 {
		return Product, errors.New("data tidak ditemukan")
	}

	return Product, nil
}

func GetCheckoutProduct () ([]Product,error) {
	var Product []Product

	err := Database.Database.Raw("SELECT * FROM products WHERE active = ? ", "true").Scan(&Product)
	if err.Error != nil {
		return Product, err.Error
	}

	if err.RowsAffected == 0 {
		return Product, errors.New("data tidak ditemukan")
	}

	return Product, nil
}

func UpdateStock (id uint, stock int) error {
	err := Database.Database.Exec("UPDATE products SET stock = ? WHERE id = ?", stock, id)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
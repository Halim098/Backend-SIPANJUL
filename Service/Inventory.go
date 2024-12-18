package Service

import (
	"Sipanjul/Controller"
	"Sipanjul/Model"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func AddProduct(c *gin.Context) {
	id := c.MustGet("id").(uint)

	data := Model.Product{}

	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"status":"fail","massage" : err.Error()})
		return
	}

	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	data.OprID = id

	err = Controller.AddProduct(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status":"fail","message": err.Error()})
		return
	}

	massage := fmt.Sprintf("Product %s berhasil ditambahkan",data.Name)
	c.JSON(http.StatusOK, gin.H{"status":"success","massage":massage})
}

func UpdateProduct(c *gin.Context) {
	id := c.MustGet("id").(uint)
	prodid,err := strconv.ParseUint(c.Param("id"),10,64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status":"fail","message": "Not Found"})
		return
	}
	prod_id := uint(prodid)
	
	prod,err := Controller.GetProductByID(prod_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status":"fail","message": "Gagal Update, Produk Tidak Ditemukan"})
		return
	}

	if id != prod.OprID {
		c.JSON(http.StatusForbidden, gin.H{"status":"fail","message": "Access Denied"})
		return
	}

	data := Model.ProductUpdate{}
	err = c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":"fail","message":"Gagal Update, Cek Kembali data"})
		return
	}

	newdata := Model.Product{
		Name: data.Name,
		Price: data.Price,
		Type: data.Type,
		Packagesize: data.Packagesize,
		Division: data.Division,
		Imageurl: data.Imageurl,
	}

	err = Controller.UpdateProduct(&prod, &newdata)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status":"fail","message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status":"success","message": fmt.Sprintf("Data %s Berhasil Diupdate",data.Name)})
}

func DeleteProduct(c *gin.Context) {
	id := c.MustGet("id").(uint)
	prodid,err := strconv.ParseUint(c.Param("id"),10,64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status":"fail","message": "Not Found"})
		return
	}
	prod_id := uint(prodid)

	prod,err := Controller.GetProductByID(prod_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status":"fail","message": "Gagal Delete, Produk Tidak Ditemukan"})
		return
	}

	if id != prod.OprID {
		c.JSON(http.StatusForbidden, gin.H{"status":"fail","message": "Akses Ditolak"})
		return
	}

	err = Controller.DeleteProduct(prod_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status":"fail","message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status":"success","message": fmt.Sprintf("Data %s Berhasil Dihapus",prod.Name)})
}

func GetProductBYOpr(c *gin.Context)  {
	id := c.MustGet("id").(uint)

	produk,err := Controller.GetProductBYOpr(id)
	if err != nil {
		c.JSON(http.StatusNotFound,gin.H{"status":"fail","message":"Data Tidak Ada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
			"status":"success",
			"message" : produk,
		},
	)
}

func GetProductByID(c *gin.Context){
	id := c.MustGet("id").(uint)
	prodid,err := strconv.ParseUint(c.Param("id"),10,64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status":"fail","message": "Not Found"})
		return
	} 
	prod_id := uint(prodid)

	product,err := Controller.GetProductByID(prod_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status":"fail","message": err.Error()})
		return
	}

	if product.OprID != id {
		c.JSON(http.StatusForbidden, gin.H{"status":"fail","message": "Akses Ditolak"})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"status":"success",
		"message": product,
	})
}

func GetAllProduct(c *gin.Context)  {
	data, err := Controller.GetAllProduct()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status":"fail","message": "Data Tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"status":"success",
		"message": data,
	})
}

func UpdateStock(c *gin.Context) {
	id := c.MustGet("id").(uint)
	prodid,err := strconv.ParseUint(c.Param("id"),10,64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status":"fail","message": "Not Found"})
		return
	}
	prod_id := uint(prodid)

	prod,err := Controller.GetProductByID(prod_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status":"fail","message": "Gagal Update, Produk Tidak Ditemukan"})
		return
	}

	if id != prod.OprID {
		c.JSON(http.StatusForbidden, gin.H{"status":"fail","message": "Access Denied"})
		return
	}

	var data map[string]interface{}
	err = c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":"fail","message":"Gagal Update, Cek Kembali data"})
		return
	}

	stock := data["stock"].(float64)
	description := data["desc"].(string)
	isNegative := data["isNegative"].(bool)

	// float64 to int
	stocknew := int(stock)

	if isNegative {
		stocknew = stocknew * -1
	}

	err= Controller.UpdateStock(stocknew,description ,&prod)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status":"fail","message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status":"success","message": fmt.Sprintf("Berhasil Merubah Stock %s",prod.Name)})
}
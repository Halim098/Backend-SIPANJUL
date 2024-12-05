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
		c.JSON(http.StatusBadRequest,gin.H{"massage" : err.Error()})
		return
	}

	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	data.OprID = id

	err = Controller.AddProduct(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	massage := fmt.Sprintf("Product %s berhasil ditambahkan",data.Name)
	c.JSON(http.StatusOK, gin.H{"massage":massage})
}

func UpdateProduct(c *gin.Context) {
	id := c.MustGet("id").(uint)
	prodid,err := strconv.ParseUint(c.Param("id"),10,64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}
	prod_id := uint(prodid)
	
	prod,err := Controller.GetProductByID(prod_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Gagal Update, Produk Tidak Ditemukan"})
		return
	}

	if id != prod.OprID {
		c.JSON(http.StatusForbidden, gin.H{"message": "Access Denied"})
		return
	}

	data := Model.Product{}
	err = c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":"Gagal Update, Cek Kembali data"})
		return
	}

	err = Controller.UpdateProduct(&prod, &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Data %s Berhasil Diupdate",data.Name)})
}

func DeleteProduct(c *gin.Context) {
	id := c.MustGet("id").(uint)
	prodid,err := strconv.ParseUint(c.Param("id"),10,64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}
	prod_id := uint(prodid)

	prod,err := Controller.GetProductByID(prod_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Gagal Delete, Produk Tidak Ditemukan"})
		return
	}

	if id != prod.OprID {
		c.JSON(http.StatusForbidden, gin.H{"message": "Akses Ditolak"})
		return
	}

	data := Model.Product{}
	err = c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":"Gagal Delete, Cek Kembali data"})
		return
	}

	err = Controller.DeleteProduct(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Data %s Berhasil Dihapus",data.Name)})
}

func GetProductBYOpr(c *gin.Context)  {
	id := c.MustGet("id").(uint)

	produk,err := Controller.GetProductBYOpr(id)
	if err != nil {
		c.JSON(http.StatusNotFound,gin.H{"message":"Data Tidak Ada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
			"message":"Berthasil Mengambil Data",
			"data" : produk,
		},
	)
}

func GetProductByID(c *gin.Context){
	id := c.MustGet("id").(uint)
	prodid,err := strconv.ParseUint(c.Param("id"),10,64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	} 
	prod_id := uint(prodid)

	product,err := Controller.GetProductByID(prod_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if product.OprID != id {
		c.JSON(http.StatusForbidden, gin.H{"message": "Akses Ditolak"})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"message": "Sukses Mengambil Data",
		"data": product,
	})
}

func GetAllProduct(c *gin.Context)  {
	data, err := Controller.GetAllProduct()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data Tidak ditemukan"})
		return
	}

	

	c.JSON(http.StatusOK,gin.H{
		"message": "Sukses Mengambil Data",
		"data": data,
	})
}

func UpdateStock(c *gin.Context) {
	id := c.MustGet("id").(uint)
	prodid,err := strconv.ParseUint(c.Param("id"),10,64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}
	prod_id := uint(prodid)

	prod,err := Controller.GetProductByID(prod_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Gagal Update, Produk Tidak Ditemukan"})
		return
	}

	if id != prod.OprID {
		c.JSON(http.StatusForbidden, gin.H{"message": "Access Denied"})
		return
	}

	var data map[string]interface{}
	err = c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":"Gagal Update, Cek Kembali data"})
		return
	}

	stock := data["quantity"].(int)
	description := data["description"].(string)

	err= Controller.UpdateStock(stock,description ,&prod)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Berhasil Merubah Stock %s",prod.Name)})
}
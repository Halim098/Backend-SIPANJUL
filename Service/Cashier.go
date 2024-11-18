package Service

import (
	"Sipanjul/Controller"
	"Sipanjul/Model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Checkout(c *gin.Context)  {
	id := c.MustGet("id").(uint)
	data := []Model.Sales_Detail{}

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Transaksi Gagal: Check Kembali Data"})
		return
	}

	stock, err := Model.GetCheckoutProduct()
	if err != nil {
		c.JSON(http.StatusInternalServerError , gin.H{"message": "Transaksi Gagal: Gagal Mengambil data dari server"})
		return
	}

	total := 0

	for _, v := range data {
		for _, s := range stock {
			if v.ProdID == s.ID {
				if v.Quantity < s.Stock {
					reason := fmt.Sprintf("Transaksi Gagal: Stock barang %s tidak mencukupi",v.Product.Name)
					c.JSON(http.StatusBadRequest, gin.H{"message": reason})
					return
				}
			}
		}
		total += v.Total
	}

	err = Controller.Checkout(id, total, &data)
	if err!= nil {
		c.JSON(http.StatusInternalServerError , gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK,gin.H{"message": "Transaksi Berhasil"})
}
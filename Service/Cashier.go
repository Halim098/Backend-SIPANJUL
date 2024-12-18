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
	data := Model.Checkout{}
	checkout := []Model.Sales_Detail{}

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":"fail","message": "Transaksi Gagal: Check Kembali Data"})
		return
	}

	stock, err := Model.GetCheckoutProduct()
	if err != nil {
		c.JSON(http.StatusInternalServerError , gin.H{"status":"fail","message": "Transaksi Gagal: Gagal Mengambil data dari server"})
		return
	}

	for _, v := range data.Items {
		for _, s := range stock {
			if v.ID == s.ID {
				if v.Quantity > s.Stock {
					reason := fmt.Sprintf("Transaksi Gagal: Stock barang %s tidak mencukupi",v.Name)
					c.JSON(http.StatusBadRequest, gin.H{"status":"fail","message": reason})
					return
				}
				checkout = append(checkout, Model.Sales_Detail{
					SalesID: id,
					ProdID: v.ID,
					Quantity: v.Quantity,
					StockAwal: s.Stock,
					StockAkhir: s.Stock - v.Quantity,
					Total: v.Price,
				})
			}
		}
	}


	err = Controller.Checkout(id, data.TotalAmount, &checkout)
	if err!= nil {
		c.JSON(http.StatusInternalServerError , gin.H{"status":"fail","message": err.Error()})
		return
	}

	c.JSON(http.StatusOK,gin.H{"status":"success","message": "Transaksi Berhasil"})
}
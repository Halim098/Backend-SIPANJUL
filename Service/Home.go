package Service

import (
	"Sipanjul/Controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLastTransaction(c *gin.Context) {
	id := c.MustGet("id").(uint)

	data, err := Controller.GetLastTransaction(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK,gin.H{"status":"success", "message":"Berhasil Mengambil Data Transaksi", "data":data})
}

func GetBestSellingItem(c *gin.Context) {
	id := c.MustGet("id").(uint)

	weekly,monthly,weeklyerr, monthlyerr := Controller.GetBestSellingItem(id)
	if weeklyerr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": weeklyerr.Error()})
		return
	}
	if monthlyerr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": monthlyerr.Error()})
		return
	}

	c.JSON(http.StatusOK,gin.H{"status":"success", "message":"Berhasil Mengambil Data Best Selling Items", "data":gin.H{"mingguan":weekly, "bulanan":monthly}})
}
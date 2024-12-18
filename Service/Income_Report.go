package Service

import (
	"Sipanjul/Controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSalesReport(c *gin.Context) {
	id := c.MustGet("id").(uint)

	harian, mingguan, bulanan, errHarian, errMingguan, errBulanan := Controller.IncomeReport(id)
	if errHarian != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": errHarian.Error()})
		return
	}
	if errMingguan != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": errMingguan.Error()})
		return
	}
	if errBulanan != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": errBulanan.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": gin.H{
			"harian":   harian,
			"mingguan": mingguan,
			"bulanan":  bulanan,
		},
	})
}
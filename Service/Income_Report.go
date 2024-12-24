package Service

import (
	"Sipanjul/Controller"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSalesReport(c *gin.Context) {
	id := c.MustGet("id").(uint)

	harian, mingguan, bulanan, errHarian, errMingguan, errBulanan := Controller.IncomeReport(id)
	if errHarian != nil && errHarian.Error() != "data tidak ditemukan" {
		log.Println("Harian error")
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": errHarian.Error()})
		return
	}
	if errMingguan != nil && errBulanan.Error() != "data tidak ditemukan" {
		log.Println("Mingguan error")
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": errMingguan.Error()})
		return
	}
	if errBulanan != nil && errBulanan.Error() != "data tidak ditemukan" {
		log.Println("Bulanan error")
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": errBulanan.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "Berhasil Mengambil Data Rangkuman Pemjualan",
		"data": gin.H{
			"harian":   harian,
			"mingguan": mingguan,
			"bulanan":  bulanan,
		},
	})
}
package Service

import (
	"Sipanjul/Controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Statistik(c *gin.Context) {
	id := c.MustGet("id").(uint)

	harian,mingguan,bulanan,tahunan,err := Controller.Statistik(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"status":"fail","message":err.Error()})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"satatus":"success",
		"message":"Berhasil Mengambil Data Statistik",
		"data":gin.H{
			"harian":harian,
			"mingguan":mingguan,
			"bulanan":bulanan,
			"tahunan":tahunan,
		},
	})
}
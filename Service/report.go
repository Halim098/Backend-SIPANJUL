package Service

import (
	"Sipanjul/Controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetReport(c *gin.Context) {
	id := c.MustGet("id").(uint)

	var data map[string]interface{}

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	typedata := data["data"].(string)
	startdate := data["startdate"].(string)
	enddate := data["enddate"].(string)
	divisi := data["divisi"].(string)
	detail := data["detail"].(string)

	if typedata == "penjualan" {
		data, err := Controller.GetSelesReport(startdate, enddate, divisi, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data})
		return
	}

	if typedata == "perubahan stok" {
		data, err := Controller.GetProductReport(startdate, enddate, divisi, detail)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "Data Tidak Ditemukan"})
}
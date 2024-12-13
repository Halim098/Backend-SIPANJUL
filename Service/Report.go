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

func Print(c *gin.Context) {
	id := c.MustGet("id").(uint)

	var data map[string]interface{}

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	startdate := data["startdate"].(string)
	enddate := data["enddate"].(string)
	divisi := data["divisi"].(string)

	selesreport, err := Controller.GetSelesReport(startdate, enddate, divisi, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	productreport, err := Controller.GetProductReport(startdate, enddate, divisi, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	datas, dataOut,err := Controller.GenerateDataExcel(selesreport, productreport)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate Excel"})
		return
	}

	fileBytes, err := Controller.GenerateExcelPenjualan(datas, dataOut)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate Excel"})
		return
	}

	// Kirim file sebagai response
	c.Header("Content-Disposition", "attachment; filename=DataPenjualan.xlsx")
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", fileBytes)
}
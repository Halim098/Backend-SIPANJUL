package Service

import (
	"Sipanjul/Controller"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetReport(c *gin.Context) {
	id := c.MustGet("id").(uint)

	var data map[string]interface{}

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":"fail","message": err.Error()})
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
			c.JSON(http.StatusInternalServerError, gin.H{"status":"fail","message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status":"success","message": data})
		return
	}

	if typedata == "perubahan" {
		newend,err := time.Parse("2006-01-02", enddate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status":"fail","message": err.Error()})
			return
		}

		newend = newend.AddDate(0, 0, 1)
		enddate = newend.Format("2006-01-02")

		data, err := Controller.GetProductReport(startdate, enddate, divisi, detail)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status":"fail","message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status":"success","message": data})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"status":"fail","message": "Data Tidak Ditemukan"})
}

func Print(c *gin.Context) {
	id := c.MustGet("id").(uint)

	var data map[string]interface{}

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":"fail","message": err.Error()})
		return
	}

	startdate := data["startdate"].(string)
	enddate := data["enddate"].(string)
	divisi := data["divisi"].(string)

	selesreport, err := Controller.GetSelesReport(startdate, enddate, divisi, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status":"fail","message": err.Error()})
		return
	}

	newdate, err := time.Parse("2006-01-02", enddate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status":"fail","message": err.Error()})
		return
	}

	newdate = newdate.AddDate(0, 0, 1)
	enddate2 := newdate.Format("2006-01-02")
	productreport, err := Controller.GetProductReport(startdate, enddate2, divisi, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status":"fail","message": err.Error()})
		return
	}

	datas, dataOut,err := Controller.GenerateDataExcel(selesreport, productreport)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status":"fail","error": "failed to generate Excel"})
		return
	}

	fileBytes, err := Controller.GenerateExcelPenjualan(datas, dataOut)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status":"fail","error": "failed to generate Excel"})
		return
	}

	// Kirim file sebagai response
	c.Header("Content-Disposition", "attachment; filename=DataPenjualan.xlsx")
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", fileBytes)
}
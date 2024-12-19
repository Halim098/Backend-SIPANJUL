package Service

import (
	"Sipanjul/Controller"
	"fmt"
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
        c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid JSON input"})
        return
    }

    startdate, ok := data["startdate"].(string)
    if !ok || startdate == "" {
        c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid or missing 'startdate'"})
        return
    }

    enddate, ok := data["enddate"].(string)
    if !ok || enddate == "" {
        c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid or missing 'enddate'"})
        return
    }

    divisi, ok := data["divisi"].(string)
    if !ok  {
        c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid or missing 'divisi'"})
        return
    }

    selesreport, err := Controller.GetSelesReport(startdate, enddate, divisi, id)
    if err != nil {

        c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
        return
    }

	newend,err := time.Parse("2006-01-02", enddate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status":"fail","message": err.Error()})
		return
	}

	newend = newend.AddDate(0, 0, 1)
	enddate = newend.Format("2006-01-02")

	productreport, err := Controller.GetProductReport(startdate, enddate, divisi, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

    datas, dataOut, err := Controller.GenerateDataExcel(selesreport, productreport)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to generate Excel data"})
        return
    }

	for _, v := range datas {
		fmt.Println("komoditi",v.Komoditi)
		fmt.Println("stockawal",v.StockAwal)
		fmt.Println("terjual",v.Terjual)
		fmt.Println("hasil",v.Hasil)
		fmt.Println("kemasan",v.Kemasan)
		fmt.Println("harga",v.Harga)
		fmt.Println("stocktambahan",v.StockTambahan)
		fmt.Println("sisa",v.Sisa)
		fmt.Println("hasil",v.Hasil)
	}

	for _, v := range dataOut {
		fmt.Println("komoditi",v.Komoditi)
		fmt.Println("deskripsi",v.Deskripsi)
		fmt.Println("jumlah",v.Jumlah)
	}

    fileBytes, err := Controller.GenerateExcelPenjualan(datas, dataOut)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to generate Excel file"})
        return
    }

    if len(fileBytes) == 0 {
        c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Generated file is empty"})
        return
    }

    filename := fmt.Sprintf("DataPenjualan_%s_to_%s.xlsx", startdate, enddate)
    c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
    c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", fileBytes)
}


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

	c.JSON(http.StatusOK,gin.H{"status":"success", "message":data})
}

func GetBestSellingItem(c *gin.Context) {
	id := c.MustGet("id").(uint)

	data, err := Controller.GetBestSellingItem(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK,gin.H{"status":"success", "message":data})
}
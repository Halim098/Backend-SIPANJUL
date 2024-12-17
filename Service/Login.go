package Service

import (
	"Sipanjul/Controller"
	"Sipanjul/Helper"
	"Sipanjul/Model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context)  {
	data := Model.OperatorLogin{}

	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":"fail","message":"Periksa Kembali Data"})
		return
	} 

	id, err := Controller.Login(data.Name, data.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status":"fail","message":"Username atau Password Salah"})
		return
	}

	token,err := Helper.GenerateJWT(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status":"fail","message":err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":"success", 
		"message": map[string]string{"token": token},
	})
}

func Register(c *gin.Context)  {
	data := Model.Operator{}

	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":"fail","message": "Periksa Kembali Data"})
		return
	}

	if data.Name == "" || data.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status":"fail","message":"Data Tidak Boleh Kosong"})
		return
	}

	err = Controller.Register(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status":"fail","message":err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status":"success","message":"Akun Berhasil Dibuat"})
}
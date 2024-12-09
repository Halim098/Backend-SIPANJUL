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
		c.JSON(http.StatusBadRequest, gin.H{"message":"Periksa Kembali Data"})
		return
	} 

	id, err := Controller.Login(data.Name, data.Password)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message":"Username atau Password Salah"})
		return
	}

	token,err := Helper.GenerateJWT(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message":err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message":"Login Berhasil", 
		"data": map[string]string{"token": token},
	})
}

func Register(c *gin.Context)  {
	data := Model.Operator{}

	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Periksa Kembali Data"})
		return
	}

	if data.Name == "" || data.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message":"Data Tidak Boleh Kosong"})
		return
	}

	err = Controller.Register(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message":err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message":"Akun Berhasil Dibuat"})
}
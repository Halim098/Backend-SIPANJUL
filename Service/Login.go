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
		"message":"Login Berhasil", 
		"data": map[string]string{"token": token},
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

func VerifyToken(c *gin.Context)  {
	id := c.MustGet("id").(uint)

	name,err := Controller.VerifyToken(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status":"fail","message":err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status":"success","message":name})
}

func GetStoreStatus(c *gin.Context)  {
	id := c.MustGet("id").(uint)

	data,err := Controller.GetStatusStore(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status":"fail","message":err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"statua":"success","storestatus":data})
}

func UpdateStoreStatus(c *gin.Context)  {
	id := c.MustGet("id").(uint)
	data := Model.StatusStore{}
	
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":"fail","message":"Periksa Kembali Data"})
		return
	}

	err = Controller.UpdateStatusStore(id, data.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status":"fail","message":err.Error()})
		return
	}

	if  !data.Status {
		message := "Toko Berhasil Ditutup"
		c.JSON(http.StatusOK, gin.H{"status":"success","message":message})
		return
	}

	message := "Toko Berhasil Dibuka"
	c.JSON(http.StatusOK, gin.H{"status":"success","message":message})
}
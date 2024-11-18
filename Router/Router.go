package Router

import (
	"Sipanjul/Middleware"
	"Sipanjul/Service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine  {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:  []string{"Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	r.POST("/login", Service.Login)
	r.POST("/register", Service.Register)
	
	// Landing Page
	r.GET("/Inventory",Service.GetAllProduct)

	v1 := r.Group("/v1", Middleware.Auth())
	{
		// Inventory
		v1.GET("/Inventory",Service.GetProductBYOpr)
		v1.GET("/Inventory/:id", Service.GetProductByID)
		v1.POST("/Inventory",Service.AddProduct)
		v1.PUT("/Inventory/:id",Service.UpdateProduct)
		v1.DELETE("/Inventory",Service.DeleteProduct)

		// Cashier
		v1.POST("/Checkout",Service.Checkout)
	}

	return r
}
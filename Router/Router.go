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
	r.GET("/inventory",Service.GetAllProduct)

	v1 := r.Group("/opr", Middleware.Auth())
	{
		// Inventory
		v1.GET("/inventory",Service.GetProductBYOpr)
		v1.GET("/inventory/:id", Service.GetProductByID)
		v1.POST("/inventory",Service.AddProduct)
		v1.PUT("/inventory/:id",Service.UpdateProduct)
		v1.DELETE("/inventory",Service.DeleteProduct)
		v1.POST("/stock/:id",Service.UpdateStock)
		v1.POST("/report",Service.GetReport)
		v1.POST("/Print",Service.Print)

		// Cashier
		v1.POST("/checkout",Service.Checkout)
	}

	return r
}
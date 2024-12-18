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
	
	//guest
	r.GET("/product",Service.GetAllProduct)
	r.GET("/store-status",Service.GetStoreStatus)
	r.GET("/verify-token",Service.VerifyToken)

	// Operator
	v1 := r.Group("/opr", Middleware.Auth())
	{
		//Home
		v1.POST("/store-status",Service.UpdateStoreStatus)
		v1.GET("/sales-report",Service.GetSalesReport)
		v1.GET("/recent-transaction",Service.GetLastTransaction)
		v1.GET("/bestselling-product",Service.GetBestSellingItem)

		// Cashier
		v1.POST("/checkout",Service.Checkout)

		// Inventory
		v1.GET("/product",Service.GetProductBYOpr)
		v1.GET("/product/:id", Service.GetProductByID) // ?????????????
		v1.POST("/product",Service.AddProduct)
		v1.PUT("/product/:id",Service.UpdateProduct)
		v1.PUT("/product/update-stock/:id",Service.UpdateStock)
		v1.DELETE("/product",Service.DeleteProduct)


		// Report
		v1.POST("/report",Service.GetReport)
		v1.POST("/print-report",Service.Print)

		// Statistik
		v1.GET("/statistik",Service.Statistik)

		// Cashier
		v1.POST("/sales-statistic",Service.Checkout)
	}

	return r
}
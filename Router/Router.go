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

	r.POST("/login", Service.Login) // done
	r.POST("/register", Service.Register) // done
	
	//guest
	r.GET("/product",Service.GetAllProduct)
	r.GET("/store-status",Service.GetStoreStatus)
	r.GET("/verify-token",Service.VerifyToken)

	// Operator
	v1 := r.Group("/opr", Middleware.Auth())
	{
		//Home
		v1.POST("/store-status",Service.UpdateStoreStatus) 
		v1.GET("/sales-report",Service.GetSalesReport) // done
		v1.GET("/recent-transaction",Service.GetLastTransaction) 
		v1.GET("/bestselling-product",Service.GetBestSellingItem) 

		// Cashier
		v1.POST("/checkout",Service.Checkout) // done

		// Inventory
		v1.GET("/product",Service.GetProductBYOpr) // done
		v1.GET("/product/:id", Service.GetProductByID) // ?????????????
		v1.POST("/product",Service.AddProduct) // done
		v1.PUT("/product/:id",Service.UpdateProduct) // done
		v1.PUT("/product/update-stock/:id",Service.UpdateStock) // done
		v1.DELETE("/product/:id",Service.DeleteProduct) // done


		// Report
		v1.POST("/report",Service.GetReport) // done
		v1.POST("/print-report",Service.Print) // done

		// Statistik
		v1.GET("/sales-statistic",Service.Statistik) // done
	}

	return r
}
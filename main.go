package handler

import (
	"Sipanjul/Database"
	"Sipanjul/Router"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	// Add logger package for GORM
)

func Handler(w http.ResponseWriter, r *http.Request) {
    loadEnv()
    Database.Connect()
	// Database.Database.AutoMigrate(&Model.Product{}, &Model.Operator{},  &Model.Sales{}, &Model.Sales_Detail{})

	// Router.SetupRouter().Run("0.0.0.0:8050")
	router := Router.SetupRouter()
	router.ServeHTTP(w, r)
}

// func main() {
//     loadEnv()
//     Database.Connect()
// 	// Database.Database.AutoMigrate(&Model.Product{}, &Model.Operator{},  &Model.Sales{}, &Model.Sales_Detail{})

// 	Router.SetupRouter().Run("0.0.0.0:8050")
// 	// router := Router.SetupRouter()
// 	// router.ServeHTTP(w, r)
// }

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}
}
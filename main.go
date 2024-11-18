package main

import (
	"Sipanjul/Database"
	"Sipanjul/Router"
	"log"

	"github.com/joho/godotenv"
	// Add logger package for GORM
)

func main() {
    loadEnv()
    Database.Connect()
	// Database.Database.AutoMigrate(&Model.Product{}, &Model.Operator{},  &Model.Sales{}, &Model.Sales_Detail{})

	Router.SetupRouter().Run("0.0.0.0:8050")
}

func loadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}
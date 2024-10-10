package main

import (
	"log"
	"Sipanjul/Database"

	"github.com/joho/godotenv"
	// Add logger package for GORM
)

func main() {
    loadEnv()
    Database.Connect()
	// Database.Database.AutoMigrate(&Model.Product{}, &Model.Operator{}, &Model.Product_Report{}, &Model.Sales{}, &Model.Sales_Detail{})
}

func loadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}
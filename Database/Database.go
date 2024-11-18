package Database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var Database *gorm.DB

func Connect() {
    host := os.Getenv("DB_HOST")
    username := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    databaseName := os.Getenv("DB_NAME")
    port := os.Getenv("DB_PORT")

    if host == "" || username == "" || password == "" || databaseName == "" || port == "" {
        panic("Database configuration is missing required environment variables")
    }

    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos",
        host, username, password, databaseName, port,
    )

    var err error
    Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database: " + err.Error())
    }

    fmt.Println("Database connection established")
}

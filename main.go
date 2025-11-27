package main

import (
    "fmt"
    "log"

    "github.com/joho/godotenv"
    "uas-backend-go/database"
)

func main() {

    // Load .env di main (sesuai clean architecture)
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env:", err)
    }

    // Connect PostgreSQL
    if err := database.InitPostgres(); err != nil {
        log.Fatal("PostgreSQL error:", err)
    }

    // Connect MongoDB
    if err := database.InitMongo(); err != nil {
        log.Fatal("MongoDB error:", err)
    }

    fmt.Println("All databases connected successfully!")
}

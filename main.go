package main

import (
    "fmt"
    "log"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "uas-backend-go/database"
    "uas-backend-go/app/repository"
    "uas-backend-go/app/service"
    "uas-backend-go/helper"
    "uas-backend-go/route"
)

func main() {

    // Load .env
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

    // ================== CLEAN WARNING ===================
    gin.SetMode(gin.ReleaseMode)

    r := gin.New()          
    r.Use(gin.Logger())     
    r.Use(gin.Recovery())   

    r.SetTrustedProxies(nil) 

  // Repository
userRepo := repository.NewUserRepository(database.Postgres)
permRepo := repository.NewPermissionRepository(database.Postgres)

// Service
authService := service.NewAuthService(userRepo, permRepo)
userService := service.NewUserService(userRepo)

// Helper
authHelper := helper.NewAuthHelper(authService)
userHelper := helper.NewUserHelper(userService)

// Routes
route.AuthRoutes(r, authHelper)
route.UserRoutes(r, userHelper) // buat route users


    r.Run(":8080")
}

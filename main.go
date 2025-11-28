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

    r := gin.New()          // tanpa default warning
    r.Use(gin.Logger())     // tambah logger manual
    r.Use(gin.Recovery())   // tambah recovery manual

    r.SetTrustedProxies(nil) // hilangkan “trusted all proxies”
    // =====================================================

    // =============== REPOSITORY ===============
    userRepo := repository.NewUserRepository(database.Postgres)
    permRepo := repository.NewPermissionRepository(database.Postgres)

    // =============== SERVICE ===============
    authService := service.NewAuthService(userRepo, permRepo)

    // =============== HELPER ===============
    authHelper := helper.NewAuthHelper(authService)

    // =============== ROUTES ===============
    route.AuthRoutes(r, authHelper)

    r.Run(":8080")
}

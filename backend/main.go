package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "backend/middlewares"
    "backend/routes"
)

func main() {
    // Muat .env
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    // Inisialisasi router Gin
    router := gin.Default()

    // Pasang middleware logger
    router.Use(middlewares.LoggerMiddleware())

    // Daftarkan route autentikasi dan rekomendasi
    routes.AuthRoutes(router)
    routes.RecommendationRoutes(router)

    log.Println("Backend running on port 8080")
    router.Run(":8080")
}
package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"backend/controllers"
	"backend/middlewares"
	"time"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 1. Konfigurasi CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // alamat FE
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 2. Public routes (tanpa auth)
	r.POST("/api/register", controllers.Register)
	r.POST("/api/login", controllers.Login)

	// 3. Protected routes (dengan middleware Auth)
	authorized := r.Group("/api")
	authorized.Use(middlewares.AuthMiddleware())
	{
		authorized.POST("/profile", controllers.CreateOrUpdateProfile)
		authorized.GET("/profile", controllers.GetProfile)

		authorized.POST("/recommend", controllers.GenerateRecommendation)
		authorized.GET("/recommend", controllers.GetRecommendations)
	}

	return r
}
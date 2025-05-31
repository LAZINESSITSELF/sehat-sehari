package routes

import (
    "github.com/gin-gonic/gin"
    "backend/controllers"
)

func RecommendationRoutes(router *gin.Engine) {
    api := router.Group("/api")
    {
        api.POST("/recommendation", controllers.GetRecommendation)
    }
}
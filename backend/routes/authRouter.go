package routes

import (
    "github.com/gin-gonic/gin"
    "backend/controllers"
)

func AuthRoutes(router *gin.Engine) {
    auth := router.Group("/api/auth")
    {
        auth.POST("/register", controllers.RegisterUser)
        auth.POST("/login", controllers.LoginUser)
        auth.PUT("/users/:id", controllers.EditUser)
        auth.DELETE("/users/:id", controllers.DeleteUser)
    }
}
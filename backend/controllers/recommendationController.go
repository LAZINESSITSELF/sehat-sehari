package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "backend/models"
    "backend/services"
)

// GetRecommendation menerima input rekomendasi dari klien dan memanggil service untuk mendapatkan rekomendasi.
func GetRecommendation(c *gin.Context) {
    var input models.UserInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
        return
    }

    reco, err := services.CallMLService(input)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error memanggil ML service"})
        return
    }
    c.JSON(http.StatusOK, reco)
}
package controllers

import (
    "context"
    "net/http"
    "os"
    "time"

    "github.com/gin-gonic/gin"
    "backend/models"
    "backend/services"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    // "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func GenerateRecommendation(c *gin.Context) {
    // Ambil collection recommendations secara “on‐the‐fly”
    recCol := services.Client.
        Database(os.Getenv("DB_NAME")).
        Collection("recommendations")

    // Ambil collection profiles secara “on‐the‐fly”
    profileCol := services.Client.
        Database(os.Getenv("DB_NAME")).
        Collection("profiles")

    userIDHex, _ := c.Get("user_id")
    userID, _ := primitive.ObjectIDFromHex(userIDHex.(string))

    // Ambil profil pengguna
    var profile models.Profile
    err := profileCol.
        FindOne(context.Background(), bson.M{"user_id": userID}).
        Decode(&profile)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Profile not found"})
        return
    }

    // Panggil ML Service
    mlReq := services.MLRequest{
        Age:              profile.Age,
        Gender:           profile.Gender,
        HeightCm:         profile.HeightCm,
        WeightKg:         profile.WeightKg,
        ActivityLevel:    profile.ActivityLevel,
        Goal:             profile.Goal,
        ExerciseDuration: profile.ExerciseDuration,
    }
    mlRes, err := services.CallMLService(mlReq)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call ML service"})
        return
    }

    // Simpan rekomendasi ke MongoDB
    rec := models.Recommendation{
        UserID:      userID,
        Date:        primitive.NewDateTimeFromTime(time.Now()),
        CalorieNeed: mlRes.CalorieNeed,
        Macros:      mlRes.Macros,
        Menu:        mlRes.Menu,
        Activities:  mlRes.Activities,
    }
    _, err = recCol.InsertOne(context.Background(), rec)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save recommendation"})
        return
    }

    c.JSON(http.StatusOK, rec)
}

func GetRecommendations(c *gin.Context) {
    // Ambil collection recommendations secara “on‐the‐fly”
    recCol := services.Client.
        Database(os.Getenv("DB_NAME")).
        Collection("recommendations")

    userIDHex, _ := c.Get("user_id")
    userID, _ := primitive.ObjectIDFromHex(userIDHex.(string))

    filter := bson.M{"user_id": userID}
    opts := options.Find().SetSort(bson.D{{Key: "date", Value: -1}}).SetLimit(7)

    cursor, err := recCol.Find(context.Background(), filter, opts)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch recommendations"})
        return
    }
    defer cursor.Close(context.Background())

    var recs []models.Recommendation
    if err = cursor.All(context.Background(), &recs); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode recommendations"})
        return
    }
    c.JSON(http.StatusOK, recs)
}

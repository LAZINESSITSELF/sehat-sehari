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
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type ProfileInput struct {
    Age              int     `json:"age" binding:"required"`
    Gender           string  `json:"gender" binding:"required,oneof=male female"`
    HeightCm         float64 `json:"height_cm" binding:"required"`
    WeightKg         float64 `json:"weight_kg" binding:"required"`
    ActivityLevel    string  `json:"activity_level" binding:"required,oneof=sedentary light moderate active very_active"`
    Goal             string  `json:"goal" binding:"required,oneof=lose_weight maintain_weight gain_weight"`
    ExerciseDuration int     `json:"exercise_duration_min" binding:"required"`
}

func CreateOrUpdateProfile(c *gin.Context) {
    // Ambil collection secara “on‐the‐fly”
    profileCollection := services.Client.
        Database(os.Getenv("DB_NAME")).
        Collection("profiles")

    userIDHex, _ := c.Get("user_id")
    userID, _ := primitive.ObjectIDFromHex(userIDHex.(string))

    var input ProfileInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    filter := bson.M{"user_id": userID}
    update := bson.M{
        "$set": bson.M{
            "age":               input.Age,
            "gender":            input.Gender,
            "height_cm":         input.HeightCm,
            "weight_kg":         input.WeightKg,
            "activity_level":    input.ActivityLevel,
            "goal":              input.Goal,
            "exercise_duration": input.ExerciseDuration,
            "updated_at":        primitive.NewDateTimeFromTime(time.Now()),
        },
    }
    opts := options.Update().SetUpsert(true)
    _, err := profileCollection.UpdateOne(context.Background(), filter, update, opts)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save profile"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Profile saved"})
}

func GetProfile(c *gin.Context) {
    // Ambil collection secara “on‐the‐fly”
    profileCollection := services.Client.
        Database(os.Getenv("DB_NAME")).
        Collection("profiles")

    userIDHex, _ := c.Get("user_id")
    userID, _ := primitive.ObjectIDFromHex(userIDHex.(string))

    var profile models.Profile
    err := profileCollection.
        FindOne(context.Background(), bson.M{"user_id": userID}).
        Decode(&profile)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch profile"})
        return
    }
    c.JSON(http.StatusOK, profile)
}

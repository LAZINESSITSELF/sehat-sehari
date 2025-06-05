package services

import (
    "context"
    "fmt"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "github.com/joho/godotenv"
    "os"
)

var Client *mongo.Client

func InitMongoDB() {
    // Load .env
    err := godotenv.Load()
    if err != nil {
        log.Println("Error loading .env file, proceeding with system env")
    }

    uri := os.Getenv("MONGODB_URI")
    if uri == "" {
        log.Fatal("MONGODB_URI is required")
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal("Cannot connect to MongoDB:", err)
    }

    // Cek koneksi
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal("MongoDB ping failed:", err)
    }
    fmt.Println("Connected to MongoDB")

    Client = client
}

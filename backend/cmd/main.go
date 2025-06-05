package main

import (
    "backend/routes"
    "backend/services"
    "log"
    "os"
)

func main() {
    // Inisialisasi Mongo & ML Client
    services.InitMongoDB()
    services.InitMLClient()

    r := routes.SetupRouter()
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Printf("Server berjalan di port %s", port)
    if err := r.Run(":" + port); err != nil {
        log.Fatal("Gagal menjalankan server:", err)
    }
}

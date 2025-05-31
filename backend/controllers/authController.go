package controllers

import (
    "net/http"
    "os"
    "strconv"
    "sync"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
    "golang.org/x/crypto/bcrypt"
    "backend/models"
)

var (
    users      = make(map[int]models.User)
    lastUserID = 0
    mu         sync.Mutex
)

func RegisterUser(c *gin.Context) {
    var input struct {
        Username string `json:"username" binding:"required"`
        Email    string `json:"email" binding:"required"`
        Password string `json:"password" binding:"required"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Tidak dapat meng-hash password"})
        return
    }

    mu.Lock()
    lastUserID++
    user := models.User{
        ID:        lastUserID,
        Username:  input.Username,
        Email:     input.Email,
        Password:  string(hashedPassword),
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }
    users[user.ID] = user
    mu.Unlock()

    c.JSON(http.StatusOK, gin.H{"message": "Registrasi berhasil", "user": user})
}

func LoginUser(c *gin.Context) {
    var input struct {
        Email    string `json:"email" binding:"required"`
        Password string `json:"password" binding:"required"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
        return
    }

    var user models.User
    found := false
    mu.Lock()
    for _, u := range users {
        if u.Email == input.Email {
            user = u
            found = true
            break
        }
    }
    mu.Unlock()
    if !found {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Kredensial tidak valid"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Kredensial tidak valid"})
        return
    }

    token, err := generateToken(strconv.Itoa(user.ID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menggenerate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Login berhasil", "token": token})
}

func EditUser(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID user tidak valid"})
        return
    }

    var input struct {
        Username string `json:"username"`
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
        return
    }

    mu.Lock()
    user, exists := users[id]
    if !exists {
        mu.Unlock()
        c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
        return
    }
    if input.Username != "" {
        user.Username = input.Username
    }
    if input.Email != "" {
        user.Email = input.Email
    }
    if input.Password != "" {
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
        if err != nil {
            mu.Unlock()
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update password"})
            return
        }
        user.Password = string(hashedPassword)
    }
    user.UpdatedAt = time.Now()
    users[id] = user
    mu.Unlock()

    c.JSON(http.StatusOK, gin.H{"message": "User berhasil diupdate", "user": user})
}

func DeleteUser(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID user tidak valid"})
        return
    }

    mu.Lock()
    _, exists := users[id]
    if !exists {
        mu.Unlock()
        c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
        return
    }
    delete(users, id)
    mu.Unlock()

    c.JSON(http.StatusOK, gin.H{"message": "User berhasil dihapus"})
}

func generateToken(userID string) (string, error) {
    jwtSecret := []byte(os.Getenv("JWT_SECRET"))
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}
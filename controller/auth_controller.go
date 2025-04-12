package controllers

import (
    "net/http"
    "github.com/salmanj7/activity-booking/config"
    "github.com/salmanj7/activity-booking/model"
    "github.com/salmanj7/activity-booking/utils"

    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
    var input struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    var user model.User
    if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    token, _ := utils.GenerateJWT(user.ID, user.Role)
    c.JSON(http.StatusOK, gin.H{"token": token})
}

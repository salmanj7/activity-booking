package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/salmanj7/activity-booking/model"
    "github.com/salmanj7/activity-booking/config"
)

func GetAllUsers(c *gin.Context) {
    var users []models.User
    config.DB.Find(&users)
    c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
    var user model.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&user)
    c.JSON(http.StatusCreated, user)
}

func GetProfile(c *gin.Context) {
    userID := c.MustGet("user_id").(uint)
    c.JSON(200, gin.H{"message": "You are logged in", "user_id": userID})
}

package controllers

import (
    "github.com/gin-gonic/gin"
    "myapp/models"
    "myapp/services"
    "net/http"
)

func Login(c *gin.Context) {
    var credentials models.UserCredentials
    if err := c.ShouldBindJSON(&credentials); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    token, err := services.Login(credentials)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

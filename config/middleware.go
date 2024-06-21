package config

import (
    "github.com/gin-gonic/gin"
    "myapp/models"
    "net/http"
    "strings"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
            c.Abort()
            return
        }

        tokenString := strings.Split(authHeader, " ")[1]
        token, err := models.VerifyToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Set("userID", token.UserID)
        c.Set("role", token.Role)
        c.Next()
    }
}

func AuthorizeRoles(roles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        userRole, _ := c.Get("role")
        authorized := false
        for _, role := range roles {
            if userRole == role {
                authorized = true
                break
            }
        }
        if !authorized {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }
        c.Next()
    }
}

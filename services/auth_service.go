package services

import (
    "github.com/dgrijalva/jwt-go"
    "assignment_project/config"
    "assignment_project/models"
    "time"
)

func Login(credentials models.UserCredentials) (string, error) {
    user, err := models.GetUserByUsername(credentials.Username)
    if err != nil {
        return "", err
    }

    if user.Password != credentials.Password {
        return "", models.ErrInvalidCredentials
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "userID": user.ID,
        "role":   user.Role,
        "exp":    time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, err := token.SignedString([]byte(config.JWTSecret))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

package models

import (
    "github.com/dgrijalva/jwt-go"
    "myapp/config"
    "errors"
)

type TokenClaims struct {
    UserID int  `json:"userID"`
    Role   Role `json:"role"`
    jwt.StandardClaims
}

func VerifyToken(tokenString string) (*TokenClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(config.JWTSecret), nil
    })
    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*TokenClaims)
    if !ok || !token.Valid {
        return nil, ErrInvalidToken
    }

    return claims, nil
}

var ErrInvalidToken = errors.New("invalid token")

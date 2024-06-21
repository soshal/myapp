package config

import (
    "database/sql"
    "log"
    "os"

    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)

var DB *sql.DB
var JWTSecret string

func Init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    JWTSecret = os.Getenv("JWT_SECRET")

    psqlInfo := os.Getenv("DATABASE_URL")
    DB, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatal(err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Successfully connected to database")
}

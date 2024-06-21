package models

import (
    "database/sql"
    "errors"
    "myapp/config"
)

var (
    ErrInvalidCredentials = errors.New("invalid credentials")
)

type Role string

const (
    RoleSales         Role = "Sales"
    RoleAccountant    Role = "Accountant"
    RoleHR            Role = "HR"
    RoleAdministrator Role = "Administrator"
)

type UserCredentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
    Role     Role   `json:"role"`
}

func GetUserByUsername(username string) (User, error) {
    var user User
    query := "SELECT id, username, password, role FROM users WHERE username=$1"
    row := config.DB.QueryRow(query, username)
    err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role)
    return user, err
}

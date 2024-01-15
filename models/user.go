package models

import "github.com/google/uuid"

type User struct {
    ID string
    Username string
    Password string
}

func NewUser(username, password string) *User {
    return &User{
        ID: uuid.New().String(),
        Username: username,
        Password: password,
    }
}

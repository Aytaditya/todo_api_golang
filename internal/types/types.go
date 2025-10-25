package types

import (
	"github.com/golang-jwt/jwt/v5"
)

type Signup struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CustomClaims struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type CreateTodo struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

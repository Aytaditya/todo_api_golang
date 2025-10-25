package jwt

import (
	"fmt"
	"time"

	"github.com/Aytaditya/todo_api_golang/internal/types"
	"github.com/golang-jwt/jwt/v5"
)

var jwt_secret = []byte("AdityaIsGoodBoy")

// generate token
func GenerateToken(userId int64, email string) (string, error) {
	claims := types.CustomClaims{
		ID:    userId,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 1 day expiry
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-app",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // creates a new token using the HS256 signing algorithm.
	return token.SignedString(jwt_secret)                      // signs the token using your secret key and returns it as a string
}

func ValidateToken(tokenString string) (*types.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &types.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwt_secret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*types.CustomClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}
	fmt.Println(claims)
	return claims, nil
}

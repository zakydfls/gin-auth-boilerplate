package entity

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

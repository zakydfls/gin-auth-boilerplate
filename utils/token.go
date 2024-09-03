package utils

import (
	"errors"
	"gin-auth-boilerplate/model/entity"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(payload entity.CustomClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateRefreshToken(username string, secret string) (string, error) {
	refreshClaims := &entity.CustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	}
	return GenerateToken(*refreshClaims, secret)

}

func VerifyToken(tokenString string, secret string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &entity.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}

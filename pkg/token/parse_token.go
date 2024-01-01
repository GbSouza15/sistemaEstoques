package token

import (
	"os"

	"github.com/GbSouza15/sistemaEstoque/internal/app/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func ParseToken(tokenString string) (uuid.UUID, error) {
	godotenv.Load()
	secret := os.Getenv("SECRET_KEY")

	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	claims, ok := token.Claims.(*models.Claims)
	if !ok {
		return uuid.Nil, err
	}

	return claims.CompanyId, nil
}

package token

import (
	"os"
	"time"

	"github.com/GbSouza15/sistemaEstoque/internal/app/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func GenerateToken(companyId uuid.UUID) (string, error) {
	godotenv.Load()
	secret := os.Getenv("SECRET_KEY")
	var company models.Company

	company.Id = companyId

	claims := &models.Claims{CompanyId: company.Id, RegisteredClaims: jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

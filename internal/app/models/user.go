package models

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	CPF       string    `json:"cpf"`
	CompanyId string    `json:"company_id"`
	Admin     bool      `json:"admin"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	CompanyId uuid.UUID `json:"company_id"`
	jwt.RegisteredClaims
}

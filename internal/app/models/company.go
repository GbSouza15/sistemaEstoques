package models

import "github.com/google/uuid"

type Company struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Address  string    `json:"address"`
	Phone    string    `json:"phone"`
	CNPJ     string    `json:"cnpj"`
}

type CompanyInfos struct {
	Id      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Address string    `json:"address"`
	Phone   string    `json:"phone"`
	CNPJ    string    `json:"cnpj"`
}

package models

import "github.com/google/uuid"

type Deposit struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CompanyId uuid.UUID `json:"company_id"`
}

type DepositUpdate struct {
	Id      uuid.UUID `json:"id"`
	Name    *string   `json:"name"`
	Address *string   `json:"address"`
}

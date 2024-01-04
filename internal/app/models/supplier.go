package models

import "github.com/google/uuid"

type Supplier struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CompanyId uuid.UUID `json:"companyId"`
}

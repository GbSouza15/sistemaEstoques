package models

import "github.com/google/uuid"

type Product struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Segment   string    `json:"segment"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock"`
	CompanyId uuid.UUID `json:"company_id"`
}

type ProductUpdate struct {
	Id      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Segment string    `json:"segment"`
	Price   float64   `json:"price"`
	Stock   int       `json:"stock"`
}

package models

import "github.com/google/uuid"

type Product struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CompanyId uuid.UUID `json:"company_id"`
	SegmentId uuid.UUID `json:"segment_id"`
}

type ProductUpdate struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	SegmentId uuid.UUID `json:"segment_id"`
}

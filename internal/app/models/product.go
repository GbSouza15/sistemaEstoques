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
	Id        string   `json:"id"`
	Name      *string  `json:"name"`
	Price     *float64 `json:"price"`
	SegmentId *string  `json:"segment_id"`
}

type ProductSegment struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CompanyId uuid.UUID `json:"company_id"`
}

// type ProductSearch struct {
// 	Name string `json:"name"`
// }

type ProductInfo struct {
	Id            string
	Name          string
	Price         float64
	CompanyId     string
	SegmentId     string
	SegmentName   string
	StockQuantity int
}

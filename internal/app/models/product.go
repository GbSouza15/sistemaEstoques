package models

import "github.com/google/uuid"

type Product struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CompanyId uuid.UUID `json:"company_id"`
	SegmentId uuid.UUID `json:"segment_id"`
	DepositId uuid.UUID `json:"deposit_id"`
	Quantity  int       `json:"quantity"`
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

type ProductSearch struct {
	Name string `json:"name"`
}

type ProductInfo struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	CompanyId     string  `json:"company_id"`
	SegmentId     string  `json:"segment_id"`
	SegmentName   string  `json:"segment_name"`
	StockQuantity int     `json:"stock_quantity"`
}

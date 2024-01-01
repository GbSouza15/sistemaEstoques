package repository

import (
	"database/sql"

	"github.com/google/uuid"
)

type ProductRepository struct {
	Db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (p *ProductRepository) CreateProduct(id uuid.UUID, name string, price float64, companyId, segment_id uuid.UUID) error {
	_, err := p.Db.Exec("INSERT INTO dev.products (id, name, price, company_id, segment_id) VALUES ($1, $2, $3, $4, $5)", id, name, price, companyId, segment_id)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) RemoveProduct(id string) error {
	_, err := p.Db.Exec("DELETE FROM dev.products WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) UpdateProduct(id uuid.UUID, name string, segmentId uuid.UUID, price float64) error {
	_, err := p.Db.Exec("UPDATE dev.products SET "+
		"name = COALESCE(NULLIF($1, ''), name), "+
		"segment_id = COALESCE(NULLIF($2, ''), segment_id), "+
		"price = COALESCE(NULLIF($3, 0), price) "+
		"WHERE id = $4",
		name, segmentId, price, id)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) AddProductSegment(id uuid.UUID, name string, company_id uuid.UUID) error {
	_, err := p.Db.Exec("INSERT INTO dev.product_segments (id, name, company_id) VALUES ($1, $2, $3)", id, name, company_id)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) RemoveProductSegment(id uuid.UUID) error {
	_, err := p.Db.Exec("DELETE FROM dev.product_segments WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

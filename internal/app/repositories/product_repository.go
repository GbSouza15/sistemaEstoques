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

func (p *ProductRepository) CreateProduct(id string, name, segment string, price float64, stock int, companyId uuid.UUID) error {
	_, err := p.Db.Exec("INSERT INTO dev.products (id, name, segment, price, stock_quantity, company_id) VALUES ($1, $2, $3, $4, $5, $6)", id, name, segment, price, stock, companyId)
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

func (p *ProductRepository) UpdateProduct(id uuid.UUID, name, segment string, price float64, stock int) error {
	_, err := p.Db.Exec("UPDATE dev.products SET "+
		"name = COALESCE(NULLIF($1, ''), name), "+
		"segment = COALESCE(NULLIF($2, ''), segment), "+
		"price = COALESCE(NULLIF($3, 0), price), "+
		"stock_quantity = COALESCE(NULLIF($4, 0), stock_quantity) "+
		"WHERE id = $5",
		name, segment, price, stock, id)

	if err != nil {
		return err
	}

	return nil
}

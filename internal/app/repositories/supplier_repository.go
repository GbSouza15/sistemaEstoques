package repository

import (
	"database/sql"

	"github.com/google/uuid"
)

type SupplierRepository struct {
	Db *sql.DB
}

func NewSupplierRepository(db *sql.DB) *SupplierRepository {
	return &SupplierRepository{db}
}

func (s *SupplierRepository) CreateSupplier(id uuid.UUID, name, email, phone string, company_id uuid.UUID) error {
	_, err := s.Db.Exec("INSERT INTO dev.suppliers (id, name, email, phone, company_id) VALUES ($1, $2, $3, $4, $5)", id, name, email, phone, company_id)
	if err != nil {
		return err
	}

	return nil
}

package repository

import (
	"database/sql"

	"github.com/GbSouza15/sistemaEstoque/internal/app/models"
	"github.com/google/uuid"
)

type SupplierRepository struct {
	Db *sql.DB
}

func NewSupplierRepository(db *sql.DB) *SupplierRepository {
	return &SupplierRepository{db}
}

func (s *SupplierRepository) CreateSupplier(id uuid.UUID, name, email, phone string, company_id uuid.UUID) error {
	_, err := s.Db.Exec("INSERT INTO suppliers (id, name, email, phone, company_id) VALUES (?, ?, ?, ?, ?)", id, name, email, phone, company_id)
	if err != nil {
		return err
	}

	return nil
}

func (s *SupplierRepository) UpdateSupplier(id uuid.UUID, name *string, email *string, phone *string) error {
	_, err := s.Db.Exec("UPDATE suppliers SET name = ?, email = ?, phone = ? WHERE id = ?", name, email, phone, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *SupplierRepository) DeleteSupplier(id string) error {
	_, err := s.Db.Exec("DELETE FROM suppliers WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (s *SupplierRepository) GetSuppliers(id uuid.UUID) ([]models.Supplier, error) {
	rows, err := s.Db.Query("SELECT id, name, email, phone FROM suppliers WHERE company_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var suppliers []models.Supplier
	for rows.Next() {
		var supplier models.Supplier
		err := rows.Scan(&supplier.Id, &supplier.Name, &supplier.Email, &supplier.Phone)
		if err != nil {
			return nil, err
		}
		suppliers = append(suppliers, supplier)
	}

	return suppliers, nil
}

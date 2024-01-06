package repository

import (
	"database/sql"

	"github.com/google/uuid"
)

type DepositRepository struct {
	Db *sql.DB
}

func NewDepositRepository(db *sql.DB) *DepositRepository {
	return &DepositRepository{db}
}

func (d *DepositRepository) CreateDeposit(id uuid.UUID, name, address string, companyId uuid.UUID) error {
	_, err := d.Db.Exec("INSERT INTO dev.deposits (id, name, address, company_id) VALUES ($1, $2, $3, $4)", id, name, address, companyId)
	if err != nil {
		return err
	}

	return nil
}

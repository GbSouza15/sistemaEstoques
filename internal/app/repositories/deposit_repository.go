package repository

import (
	"database/sql"

	"github.com/GbSouza15/sistemaEstoque/internal/app/models"
	"github.com/google/uuid"
)

type DepositRepository struct {
	Db *sql.DB
}

func NewDepositRepository(db *sql.DB) *DepositRepository {
	return &DepositRepository{db}
}

func (d *DepositRepository) CreateDeposit(id uuid.UUID, name, address string, companyId uuid.UUID) error {
	_, err := d.Db.Exec("INSERT INTO deposits (id, name, address, company_id) VALUES (?, ?, ?, ?)", id, name, address, companyId)
	if err != nil {
		return err
	}

	return nil
}

func (d *DepositRepository) ListDeposits(companyId uuid.UUID) ([]models.Deposit, error) {
	rows, err := d.Db.Query("SELECT id, name, address FROM deposits WHERE company_id = ?", companyId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var deposits []models.Deposit
	for rows.Next() {
		var deposit models.Deposit
		err := rows.Scan(&deposit.Id, &deposit.Name, &deposit.Address)
		if err != nil {
			return nil, err
		}
		deposits = append(deposits, deposit)
	}

	return deposits, nil
}

func (d *DepositRepository) DeleteDeposit(id string) error {
	_, err := d.Db.Exec("DELETE FROM deposits WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (d *DepositRepository) UpdateDeposit(id uuid.UUID, name *string, address *string) error {
	_, err := d.Db.Exec("UPDATE deposits SET name = COALESCE(?, name), address = COALESCE(?, address) WHERE id = ?", name, address, id)
	if err != nil {
		return err
	}

	return nil
}

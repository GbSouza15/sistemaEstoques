package repository

import (
	"database/sql"

	"github.com/google/uuid"
)

type CompanyRepository struct {
	Db *sql.DB
}

func NewCompanyRepository(db *sql.DB) *CompanyRepository {
	return &CompanyRepository{db}
}

func (c *CompanyRepository) RegisterCompany(id, name, email, password, address, phone, cnpj string) error {
	_, err := c.Db.Exec("INSERT INTO companies (id, name, email, password, address, phone, cnpj) VALUES (?, ?, ?, ?, ?, ?, ?)", id, name, email, password, address, phone, cnpj)
	if err != nil {
		return err
	}

	return nil
}

func (c *CompanyRepository) GetCompanyInfos(id uuid.UUID) (string, string, string, string, string, error) {
	var name, email, address, phone, cnpj string

	err := c.Db.QueryRow("SELECT name, email, address, phone, cnpj FROM companies WHERE id = ?", id).Scan(&name, &email, &address, &phone, &cnpj)
	if err != nil {
		return "", "", "", "", "", err
	}

	return name, email, address, phone, cnpj, nil
}

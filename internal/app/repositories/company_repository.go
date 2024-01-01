package repository

import (
	"database/sql"
)

type CompanyRepository struct {
	Db *sql.DB
}

func NewCompanyRepository(db *sql.DB) *CompanyRepository {
	return &CompanyRepository{db}
}

func (c *CompanyRepository) RegisterCompany(id, name, email, password, address, phone, cnpj string) error {
	_, err := c.Db.Exec("INSERT INTO dev.companies (id, name, email, password, address, phone, cnpj) VALUES ($1, $2, $3, $4, $5, $6, $7)", id, name, email, password, address, phone, cnpj)
	if err != nil {
		return err
	}

	return nil
}

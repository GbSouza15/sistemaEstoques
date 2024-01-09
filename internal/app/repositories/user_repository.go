package repository

import (
	"database/sql"

	"github.com/google/uuid"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) RegisterUser(id string, name string, email string, password string, phone string, cpf string, companyId string, admin bool) error {
	_, err := u.Db.Exec("INSERT INTO dev.users (id, name, email, password, phone, cpf, company_id, admin) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", id, name, email, password, phone, cpf, companyId, admin)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) LoginUser(email string) (uuid.UUID, string, error) {
	var companyId uuid.UUID
	var password string

	err := u.Db.QueryRow("SELECT company_id, password FROM dev.users WHERE email = $1", email).Scan(&companyId, &password)
	if err != nil {
		return uuid.Nil, "", err
	}

	return companyId, password, nil
}

func (u *UserRepository) GetCompanyInfos(id uuid.UUID) (string, string, string, string, string, error) {
	var name, email, address, phone, cnpj string

	err := u.Db.QueryRow("SELECT name, email, address, phone, cnpj FROM dev.companies WHERE id = $1", id).Scan(&name, &email, &address, &phone, &cnpj)
	if err != nil {
		return "", "", "", "", "", err
	}

	return name, email, address, phone, cnpj, nil
}

func (u *UserRepository) DeleteUser(id string) error {
	_, err := u.Db.Exec("DELETE FROM dev.users WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

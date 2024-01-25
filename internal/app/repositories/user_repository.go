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
	_, err := u.Db.Exec("INSERT INTO users (id, name, email, password, phone, cpf, company_id, admin) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", id, name, email, password, phone, cpf, companyId, admin)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) LoginUser(email string) (uuid.UUID, string, error) {
	var companyId uuid.UUID
	var password string

	err := u.Db.QueryRow("SELECT company_id, password FROM users WHERE email = ?", email).Scan(&companyId, &password)
	if err != nil {
		return uuid.Nil, "", err
	}

	return companyId, password, nil
}

func (u *UserRepository) DeleteUser(id string) error {
	_, err := u.Db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

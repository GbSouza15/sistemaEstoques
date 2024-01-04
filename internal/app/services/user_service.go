package service

import (
	"encoding/json"

	"github.com/GbSouza15/sistemaEstoque/internal/app/models"
	repository "github.com/GbSouza15/sistemaEstoque/internal/app/repositories"
	"github.com/GbSouza15/sistemaEstoque/pkg/token"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (us *UserService) RegisterUser(body []byte) error {

	var newUser models.User
	var userId = uuid.NewString()

	if err := json.Unmarshal(body, &newUser); err != nil {
		return err
	}

	hashPasswordUser, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)
	if err != nil {
		return err
	}

	if err := us.repo.RegisterUser(userId, newUser.Name, newUser.Email, string(hashPasswordUser), newUser.Phone, newUser.CPF, newUser.CompanyId, newUser.Admin); err != nil {
		return err
	}

	return nil
}

func (us *UserService) LoginUser(body []byte) (string, error) {

	var userLogin models.UserLogin

	if err := json.Unmarshal(body, &userLogin); err != nil {
		return "", err
	}

	companyId, password, err := us.repo.LoginUser(userLogin.Email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(userLogin.Password)); err != nil {
		return "", err
	}

	tokenString, err := token.GenerateToken(companyId)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (us *UserService) GetCompanyInfos(companyId uuid.UUID) (models.CompanyInfos, error) {

	companyName, companyEmail, companyAddress, companyPhone, companyCnpj, err := us.repo.GetCompanyInfos(companyId)
	if err != nil {
		return models.CompanyInfos{}, err
	}

	company := models.CompanyInfos{
		Name:    companyName,
		Email:   companyEmail,
		Address: companyAddress,
		Phone:   companyPhone,
		CNPJ:    companyCnpj,
	}

	return company, nil
}

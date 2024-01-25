package service

import (
	"encoding/json"

	"github.com/GbSouza15/sistemaEstoque/internal/app/models"
	repository "github.com/GbSouza15/sistemaEstoque/internal/app/repositories"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CompanyService struct {
	repo *repository.CompanyRepository
}

func NewCompanyService(repo *repository.CompanyRepository) *CompanyService {
	return &CompanyService{repo}
}

func (cs *CompanyService) RegisterCompany(body []byte) error {

	var newCompany models.Company

	var companyId = uuid.NewString()

	if err := json.Unmarshal(body, &newCompany); err != nil {
		return err
	}

	hashPasswordCompany, err := bcrypt.GenerateFromPassword([]byte(newCompany.Password), 14)
	if err != nil {
		return err
	}

	if err := cs.repo.RegisterCompany(companyId, newCompany.Name, newCompany.Email, string(hashPasswordCompany), newCompany.Address, newCompany.Phone, newCompany.CNPJ); err != nil {
		return err
	}

	return nil
}

func (cs *CompanyService) GetCompanyInfos(companyId uuid.UUID) (models.CompanyInfos, error) {

	companyName, companyEmail, companyAddress, companyPhone, companyCnpj, err := cs.repo.GetCompanyInfos(companyId)
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

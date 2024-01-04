package service

import (
	"encoding/json"

	"github.com/GbSouza15/sistemaEstoque/internal/app/models"
	repository "github.com/GbSouza15/sistemaEstoque/internal/app/repositories"
	"github.com/google/uuid"
)

type SuppliersService struct {
	repo *repository.SupplierRepository
}

func NewSuppliersService(repo *repository.SupplierRepository) *SuppliersService {
	return &SuppliersService{repo}
}

func (ss *SuppliersService) CreateSupplier(body []byte) error {
	var newSupplier models.Supplier

	var supplierId, err = uuid.NewUUID()
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &newSupplier); err != nil {
		return err
	}

	if err := ss.repo.CreateSupplier(supplierId, newSupplier.Name, newSupplier.Email, newSupplier.Phone, newSupplier.CompanyId); err != nil {
		return err
	}

	return nil
}

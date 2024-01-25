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

func (ss *SuppliersService) UpdateSupplier(body []byte) error {
	var supplier models.SupplierUpdate

	if err := json.Unmarshal(body, &supplier); err != nil {
		return err
	}

	if err := ss.repo.UpdateSupplier(supplier.Id, supplier.Name, supplier.Email, supplier.Phone); err != nil {
		return err
	}

	return nil
}

func (ss *SuppliersService) DeleteSupplier(id string) error {
	if err := ss.repo.DeleteSupplier(id); err != nil {
		return err
	}

	return nil
}

func (ss *SuppliersService) GetSuppliers(id uuid.UUID) ([]byte, error) {
	suppliers, err := ss.repo.GetSuppliers(id)
	if err != nil {
		return nil, err
	}

	responseJson, err := json.Marshal(suppliers)
	if err != nil {
		return nil, err
	}

	return responseJson, nil
}

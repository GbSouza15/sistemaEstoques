package service

import (
	"encoding/json"

	"github.com/GbSouza15/sistemaEstoque/internal/app/models"
	repository "github.com/GbSouza15/sistemaEstoque/internal/app/repositories"
	"github.com/google/uuid"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo}
}

func (ps *ProductService) CreateProduct(body []byte) error {
	var newProduct models.Product

	var productId = uuid.NewString()

	if err := json.Unmarshal(body, &newProduct); err != nil {
		return err
	}

	if err := ps.repo.CreateProduct(productId, newProduct.Name, newProduct.Segment, newProduct.Price, newProduct.Stock, newProduct.CompanyId); err != nil {
		return err
	}

	return nil
}

func (ps *ProductService) RemoveProduct(id string) error {
	if err := ps.repo.RemoveProduct(id); err != nil {
		return err
	}

	return nil
}

func (ps *ProductService) UpdateProduct(body []byte) error {
	var productUpdate models.ProductUpdate

	if err := json.Unmarshal(body, &productUpdate); err != nil {
		return err
	}

	if err := ps.repo.UpdateProduct(productUpdate.Id, productUpdate.Name, productUpdate.Segment, productUpdate.Price, productUpdate.Stock); err != nil {
		return err
	}

	return nil
}

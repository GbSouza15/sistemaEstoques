package service

import (
	repository "github.com/GbSouza15/sistemaEstoque/internal/app/repositories"
)

type SuppliersService struct {
	repo *repository.SupplierRepository
}

func NewSuppliersService(repo *repository.SupplierRepository) *SuppliersService {
	return &SuppliersService{repo}
}

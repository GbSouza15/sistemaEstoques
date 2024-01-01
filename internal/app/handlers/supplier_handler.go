package handler

import (
	service "github.com/GbSouza15/sistemaEstoque/internal/app/services"
)

type SupplierHandler struct {
	service *service.SuppliersService
}

func NewSupplierHandler(service *service.SuppliersService) *SupplierHandler {
	return &SupplierHandler{service}
}

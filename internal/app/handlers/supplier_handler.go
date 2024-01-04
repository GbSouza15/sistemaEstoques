package handler

import (
	"fmt"
	"io"
	"net/http"

	service "github.com/GbSouza15/sistemaEstoque/internal/app/services"
)

type SupplierHandler struct {
	service *service.SuppliersService
}

func NewSupplierHandler(service *service.SuppliersService) *SupplierHandler {
	return &SupplierHandler{service}
}

func (sh *SupplierHandler) CreateSupplier(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	defer r.Body.Close()

	if err := sh.service.CreateSupplier(body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Supplier created successfully"))
}

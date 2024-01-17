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

// @Summary Criar um novo fornecedor
// @Description Cria um novo fornecedor no sistema
// @Tags Fornecedor
// @Accept json
// @Produce json
// @Param requestBody body models.Supplier true "Detalhes do cadastro do fornecedor"
// @Success 201 {string} string "Fornecedor criado com sucesso"
// @Failure 400 {string} string "Requisição Inválida"
// @Failure 500 {string} string "Erro Interno do Servidor"
// @Router /company/register/supplier [post]
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

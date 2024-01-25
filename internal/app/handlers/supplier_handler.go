package handler

import (
	"fmt"
	"io"
	"net/http"

	service "github.com/GbSouza15/sistemaEstoque/internal/app/services"
	"github.com/GbSouza15/sistemaEstoque/pkg/token"
	"github.com/GbSouza15/sistemaEstoque/pkg/utils"
	"github.com/gorilla/mux"
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
		utils.WriteResponse("Error reading body", w, http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	defer r.Body.Close()

	if err := sh.service.CreateSupplier(body); err != nil {
		utils.WriteResponse("Error creating supplier", w, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	utils.WriteResponse("Supplier created successfully", w, http.StatusCreated)
}

// @Summary Atualizar um fornecedor
// @Description Atualiza um fornecedor no sistema
// @Tags Fornecedor
// @Accept json
// @Produce json
// @Param requestBody body models.SupplierUpdate true "Detalhes da atualização do fornecedor"
// @Success 200 {string} string "Fornecedor atualizado com sucesso"
// @Failure 400 {string} string "Requisição Inválida"
// @Failure 500 {string} string "Erro Interno do Servidor"
// @Router /company/update/supplier [put]
func (sh *SupplierHandler) UpdateSupplier(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteResponse("Error reading body", w, http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	defer r.Body.Close()

	if err := sh.service.UpdateSupplier(body); err != nil {
		utils.WriteResponse("Error updating supplier", w, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	utils.WriteResponse("Supplier updated successfully", w, http.StatusOK)
}

// @Summary Deletar um fornecedor
// @Description Deleta um fornecedor no sistema
// @Tags Fornecedor
// @Accept json
// @Produce json
// @Param id path string true "ID do fornecedor"
// @Success 200 {string} string "Fornecedor deletado com sucesso"
// @Failure 400 {string} string "Requisição Inválida"
// @Failure 500 {string} string "Erro Interno do Servidor"
// @Router /company/delete/supplier/{id} [delete]
func (sh *SupplierHandler) DeleteSupplier(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := sh.service.DeleteSupplier(id); err != nil {
		utils.WriteResponse("Error deleting supplier", w, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	utils.WriteResponse("Supplier deleted successfully", w, http.StatusOK)
}

// @Summary Listar fornecedores
// @Description Lista todos os fornecedores de uma empresa
// @Tags Fornecedor
// @Accept json
// @Produce json
// @Param id path string true "ID da empresa"
// @Success 200 {array} models.Supplier "Lista de fornecedores"
// @Failure 400 {string} string "Requisição Inválida"
// @Failure 500 {string} string "Erro Interno do Servidor"
// @Router /company/suppliers [get]
func (sh *SupplierHandler) GetSuppliers(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		utils.WriteResponse("Error getting cookie", w, http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	companyId, err := token.ParseToken(c.Value)
	if err != nil {
		utils.WriteResponse("Error parsing token", w, http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	suppliers, err := sh.service.GetSuppliers(companyId)
	if err != nil {
		utils.WriteResponse("Error getting suppliers", w, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(suppliers)
}

package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	service "github.com/GbSouza15/sistemaEstoque/internal/app/services"
	"github.com/GbSouza15/sistemaEstoque/pkg/utils"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service}
}

// @Summary Criar um novo produto
// @Description Cria um novo produto no sistema
// @Tags Produto
// @Accept json
// @Produce json
// @Param requestBody body models.Product true "Detalhes do cadastro do produto"
// @Success 201 {string} string "Produto criado com sucesso"
// @Failure 400 {string} string "Requisição Inválida"
// @Failure 500 {string} string "Erro Interno do Servidor"
// @Router /company/register/product [post]
func (ph *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteResponse("Error reading body", w, http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	defer r.Body.Close()

	if err := ph.service.CreateProduct(body); err != nil {
		utils.WriteResponse("Error creating product", w, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	utils.WriteResponse("Product created successfully", w, http.StatusCreated)
}

// @Summary Remover um produto
// @Description Remove um produto com base no ID fornecido
// @Tags Produto
// @Param productId path string true "ID do produto a ser removido"
// @Success 200 {string} string "Produto removido com sucesso"
// @Failure 500 {string} string "Erro Interno do Servidor"
// @Router /company/remove/product/{productId} [delete]
func (ph *ProductHandler) RemoveProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["productId"]

	if err := ph.service.RemoveProduct(id); err != nil {
		utils.WriteResponse("Error to remove product", w, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	utils.WriteResponse("Product removed successfully", w, http.StatusOK)
}

// @Summary Atualizar um produto
// @Description Atualiza as informações de um produto existente
// @Tags Produto
// @Accept json
// @Produce json
// @Param requestBody body models.ProductUpdate true "Detalhes de atualização do produto"
// @Success 200 {string} string "Produto atualizado com sucesso"
// @Failure 400 {string} string "Requisição Inválida"
// @Failure 500 {string} string "Erro Interno do Servidor"
// @Router /company/update/product [put]
func (ph *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteResponse("Error reading body", w, http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	defer r.Body.Close()

	if err := ph.service.UpdateProduct(body); err != nil {
		utils.WriteResponse("Error to update product", w, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	utils.WriteResponse("Product updated successfully", w, http.StatusOK)
}

// @Summary Adicionar segmento de produto
// @Description Adiciona um novo segmento a um produto existente
// @Tags Produto
// @Accept json
// @Produce json
// @Param requestBody body models.ProductSegment true "Detalhes do segmento a ser adicionado"
// @Success 201 {string} string "Segmento de produto adicionado com sucesso"
// @Failure 400 {string} string "Requisição Inválida"
// @Failure 500 {string} string "Erro Interno do Servidor"
// @Router /company/register/product/segment [post]
func (ph *ProductHandler) AddProductSegment(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteResponse("Error reading body", w, http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	defer r.Body.Close()

	if err := ph.service.AddProductSegment(body); err != nil {
		utils.WriteResponse("Error to add product segment", w, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	utils.WriteResponse("Product segment added successfully", w, http.StatusCreated)
}

// @Summary Pesquisar produtos
// @Description Pesquisa produtos com base nos critérios fornecidos
// @Tags Produto
// @Accept json
// @Produce json
// @Param requestBody body models.ProductSearch true "Critérios de pesquisa do produto"
// @Success 200 {array} models.Product "Lista de produtos correspondentes à pesquisa"
// @Failure 400 {string} string "Requisição Inválida"
// @Failure 500 {string} string "Erro Interno do Servidor"
// @Router /company/search/product [post]
func (ph *ProductHandler) SearchProduct(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteResponse("Error reading body", w, http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	products, err := ph.service.SearchProduct(body)
	if err != nil {
		utils.WriteResponse("Error to search product", w, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	responseJson, err := json.Marshal(products)
	if err != nil {
		utils.WriteResponse("Error to marshal response", w, http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

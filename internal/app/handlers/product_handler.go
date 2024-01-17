package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	service "github.com/GbSouza15/sistemaEstoque/internal/app/services"
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if err := ph.service.CreateProduct(body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Product created successfully"))
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
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error to remove product"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product removed successfully"))
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if err := ph.service.UpdateProduct(body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product updated successfully"))
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if err := ph.service.AddProductSegment(body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Product segment created successfully"))
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	products, err := ph.service.SearchProduct(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	responseJson, err := json.Marshal(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

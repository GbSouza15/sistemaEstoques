package handler

import (
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

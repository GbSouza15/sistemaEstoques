package handler

import (
	"fmt"
	"io"
	"net/http"

	service "github.com/GbSouza15/sistemaEstoque/internal/app/services"
)

type CompanyHandler struct {
	service *service.CompanyService
}

func NewCompanyHandler(service *service.CompanyService) *CompanyHandler {
	return &CompanyHandler{service}
}

func (ch *CompanyHandler) RegisterCompany(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error reading body"))
		fmt.Println(err.Error())
		return
	}

	defer r.Body.Close()

	if err := ch.service.RegisterCompany(body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error registering company"))
		fmt.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Company registered successfully"))
}

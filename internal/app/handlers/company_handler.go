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

// @Summary Registrar uma nova empresa
// @Description Registra uma nova empresa no sistema
// @Tags Empresa
// @Accept json
// @Produce json
// @Param requestBody body models.Company true "Detalhes do cadastro da empresa"
// @Success 201 {string} string "Empresa registrada com sucesso"
// @Failure 400 {string} string "Requisição Inválida"
// @Failure 500 {string} string "Erro Interno do Servidor"
// @Router /company/register [post]
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

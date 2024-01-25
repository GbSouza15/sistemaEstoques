package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	service "github.com/GbSouza15/sistemaEstoque/internal/app/services"
	"github.com/GbSouza15/sistemaEstoque/pkg/token"
	"github.com/GbSouza15/sistemaEstoque/pkg/utils"
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
		utils.WriteResponse("Error reading body", w, http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	defer r.Body.Close()

	if err := ch.service.RegisterCompany(body); err != nil {
		utils.WriteResponse("Error creating company", w, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	utils.WriteResponse("Company created successfully", w, http.StatusCreated)
}

// @Summary Pegar informações da empresa
// @Description Pegar informações da empresa baseado no ID da empresa presente no token
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} models.Company "Company information"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /company/info [get]
func (ch *CompanyHandler) GetCompanyInfos(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		utils.WriteResponse("Error getting token", w, http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	if c == nil {
		utils.WriteResponse("No token found", w, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	companyId, err := token.ParseToken(c.Value)
	if err != nil {
		utils.WriteResponse("Error parsing token", w, http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	company, err := ch.service.GetCompanyInfos(companyId)
	if err != nil {
		utils.WriteResponse("Error getting company information", w, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	reponseJson, err := json.Marshal(company)
	if err != nil {
		utils.WriteResponse("Error parsing company information", w, http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(reponseJson)
}

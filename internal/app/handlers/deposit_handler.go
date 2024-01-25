package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	service "github.com/GbSouza15/sistemaEstoque/internal/app/services"
	"github.com/GbSouza15/sistemaEstoque/pkg/token"
	"github.com/GbSouza15/sistemaEstoque/pkg/utils"
	"github.com/gorilla/mux"
)

type DepositHandler struct {
	service *service.DepositService
}

func NewDepositHandler(service *service.DepositService) *DepositHandler {
	return &DepositHandler{service}
}

// @Summary Criar um novo depósito
// @Description Cria um novo depósito no sistema
// @Tags Depósito
// @Accept json
// @Produce json
// @Param requestBody body models.Deposit true "Detalhes do cadastro do depósito"
// @Success 201 {string} string "Depósito criado com sucesso"
// @Example {string} "Deposit created successfully"
// @Failure 400 {string} string "Requisição Inválida"
// @Example {string} "Error reading body"
// @Failure 500 {string} string "Erro Interno do Servidor"
// @Example {string} "Error creating deposit"
// @Router /company/register/deposit [post]
func (dh *DepositHandler) CreateDeposit(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteResponse("Error reading body", w, http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	defer r.Body.Close()

	if err := dh.service.CreateDeposit(body); err != nil {
		utils.WriteResponse("Error creating deposit", w, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	utils.WriteResponse("Deposit created successfully", w, http.StatusCreated)
}

//teste

// @Summary Listar depósitos
// @Description Lista todos os depósitos cadastrados no sistema
// @Tags Depósito
// @Produce json
// @Param companyId query string true "ID da empresa"
// @Success 200 {array} models.Deposit "Lista de depósitos"
// @Example {array} models.Deposit
// @Failure 400 {string} string "Invalid request"
// @Example {string} "Error getting cookie"
// @Failure 500 {string} string "Internal Server Error"
// @Example {string} "Error parsing token"
// @Router /company/list/deposit [get]
func (dh *DepositHandler) ListDeposits(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		utils.WriteResponse("Error getting cookie", w, http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	if c == nil {
		utils.WriteResponse("Cookie not found", w, http.StatusBadRequest)
		return
	}

	companyId, err := token.ParseToken(c.Value)
	if err != nil {
		utils.WriteResponse("Error parsing token", w, http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	deposits, err := dh.service.ListDeposits(companyId)
	if err != nil {
		utils.WriteResponse("Error getting deposits", w, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	reponseJson, err := json.Marshal(deposits)
	if err != nil {
		utils.WriteResponse("Error parsing deposits", w, http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(reponseJson)
}

// @Summary Deletar um depósito
// @Description Deleta um depósito do sistema
// @Tags Depósito
// @Produce json
// @Param depositId query string true "ID do depósito"
// @Success 200 {string} string "Deposit removed successfully"
// @Example {string} "Deposit removed successfully"
// @Failure 500 {string} string "Error to remove deposit"
// @Example {string} "Error to remove deposit"
// @Router /company/delete/deposit/{depositId} [delete]
func (dh *DepositHandler) DeleteDeposit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["depositId"]

	if err := dh.service.DeleteDeposit(id); err != nil {
		utils.WriteResponse("Error to remove deposit", w, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	utils.WriteResponse("Deposit removed successfully", w, http.StatusOK)
}

// @Summary Atualizar um depósito
// @Description Atualiza um depósito no sistema
// @Tags Depósito
// @Accept json
// @Produce json
// @Param requestBody body models.DepositUpdate true "Detalhes do cadastro do depósito"
// @Success 200 {string} string "Deposit updated successfully"
// @Example {string} "Deposit updated successfully"
// @Failure 400 {string} string "Error reading body"
// @Example {string} "Error reading body"
// @Failure 500 {string} string "Error updating deposit"
// @Example {string} "Error updating deposit"
// @Router /company/update/deposit [put]
func (dh *DepositHandler) UpdateDeposit(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteResponse("Error reading body", w, http.StatusInternalServerError)
		return
	}

	if err := dh.service.UpdateDeposit(body); err != nil {
		utils.WriteResponse("Error updating deposit", w, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	utils.WriteResponse("Deposit updated successfully", w, http.StatusOK)
}

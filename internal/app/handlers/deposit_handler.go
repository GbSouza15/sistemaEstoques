package handler

import (
	"io"
	"net/http"

	service "github.com/GbSouza15/sistemaEstoque/internal/app/services"
)

type DepositHandler struct {
	service *service.DepositService
}

func NewDepositHandler(service *service.DepositService) *DepositHandler {
	return &DepositHandler{service}
}

func (dh *DepositHandler) CreateDeposit(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error reading body"))
		return
	}

	defer r.Body.Close()

	if err := dh.service.CreateDeposit(body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating deposit"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Deposit created successfully"))
}

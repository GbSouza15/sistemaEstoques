package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GbSouza15/sistemaEstoque/internal/app/models"
)

func WriteResponse(message string, w http.ResponseWriter, statusCode int) {
	response := models.Response{Message: message}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}

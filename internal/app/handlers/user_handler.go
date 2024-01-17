package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	service "github.com/GbSouza15/sistemaEstoque/internal/app/services"
	"github.com/GbSouza15/sistemaEstoque/pkg/token"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service}
}

// @Summary Registrar novo usuário
// @Description Registrar um novo usuário no sistema
// @Tags User
// @Accept json
// @Produce json
// @Param requestBody body models.User true "User registration details"
// @Success 201 {string} string "User registered successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /company/register/user [post]
func (us *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error reading body"))
		fmt.Println(err.Error())
		return
	}

	defer r.Body.Close()

	if err := us.service.RegisterUser(body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error registering user"))
		fmt.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}

// @Summary Login do usuário
// @Description Login do usuário no sistema
// @Tags User
// @Accept json
// @Produce json
// @Param requestBody body models.UserLogin true "User login details"
// @Success 200 {string} string "User logged in successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /company/login/user [post]
func (us *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error reading body"))
		fmt.Println(err.Error())
		return
	}

	defer r.Body.Close()

	tokenString, err := us.service.LoginUser(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error logging in user"))
		fmt.Println(err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: time.Now().Add(time.Hour * 24),
		Path:    "/company",
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User logged in successfully"))
}

// @Summary Pegar informações da empresa
// @Description Pegar informações da empresa baseado no ID da empresa presente no token
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} models.Company "Company information"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /company/user [get]
func (us *UserHandler) GetCompanyInfos(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error getting cookie"))
		fmt.Println(err.Error())
		return
	}

	if c == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Token cookie not present"))
		return
	}

	companyId, err := token.ParseToken(c.Value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error parsing token"))
		fmt.Println(err.Error())
		return
	}

	company, err := us.service.GetCompanyInfos(companyId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error getting company infos"))
		fmt.Println(err.Error())
		return
	}

	reponseJson, err := json.Marshal(company)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error parsing json"))
		fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(reponseJson)
}

// @Summary Deletar usuário
// @Description Deletar usuário do sistema baseado no ID do usuário
// @Tags User
// @Accept json
// @Produce json
// @Param userId path string true "User ID to be deleted"
// @Success 200 {string} string "User deleted successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /company/remove/user/{userId} [delete]
func (us *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	if err := us.service.DeleteUser(userId); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error deleting user"))
		fmt.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}

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

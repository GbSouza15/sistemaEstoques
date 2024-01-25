package handler

import (
	"fmt"
	"io"
	"net/http"
	"time"

	service "github.com/GbSouza15/sistemaEstoque/internal/app/services"
	"github.com/GbSouza15/sistemaEstoque/pkg/utils"
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
		utils.WriteResponse("Error reading body", w, http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	defer r.Body.Close()

	if err := us.service.RegisterUser(body); err != nil {
		utils.WriteResponse("Error registering user", w, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	utils.WriteResponse("User registered successfully", w, http.StatusCreated)
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
// @Router /company/user/login [post]
func (us *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteResponse("Error reading body", w, http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	defer r.Body.Close()

	tokenString, err := us.service.LoginUser(body)
	if err != nil {
		utils.WriteResponse("Error logging in user", w, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: time.Now().Add(time.Hour * 24),
		Path:    "/company",
	})

	utils.WriteResponse("User logged in successfully", w, http.StatusOK)
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
		utils.WriteResponse("Error deleting user", w, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	utils.WriteResponse("User deleted successfully", w, http.StatusOK)
}

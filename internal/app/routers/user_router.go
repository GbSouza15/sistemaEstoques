package router

import (
	"database/sql"
	"net/http"

	handler "github.com/GbSouza15/sistemaEstoque/internal/app/handlers"
	repository "github.com/GbSouza15/sistemaEstoque/internal/app/repositories"
	service "github.com/GbSouza15/sistemaEstoque/internal/app/services"
	"github.com/GbSouza15/sistemaEstoque/pkg/middleware"
	"github.com/gorilla/mux"
)

func UserRoutes(db *sql.DB, r *mux.Router) error {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	r.HandleFunc("/company/register/user", userHandler.RegisterUser).Methods(http.MethodPost)
	r.HandleFunc("/company/remove/user/{userId}", middleware.Middleware(userHandler.DeleteUser)).Methods(http.MethodDelete)
	r.HandleFunc("/company/user/login", userHandler.LoginUser).Methods(http.MethodPost)

	return nil
}

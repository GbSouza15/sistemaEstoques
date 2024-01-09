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

func DepositRoutes(db *sql.DB, r *mux.Router) error {
	depositRepository := repository.NewDepositRepository(db)
	depositService := service.NewDepositService(depositRepository)
	depositHandler := handler.NewDepositHandler(depositService)

	r.HandleFunc("/company/register/deposit", middleware.Middleware(depositHandler.CreateDeposit)).Methods(http.MethodPost)

	return nil
}

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

func SupplierRoutes(db *sql.DB, r *mux.Router) error {
	supplierRepository := repository.NewSupplierRepository(db)
	supplierService := service.NewSuppliersService(supplierRepository)
	supplierHandler := handler.NewSupplierHandler(supplierService)

	r.HandleFunc("/company/register/supplier", middleware.Middleware(supplierHandler.CreateSupplier)).Methods(http.MethodPost)

	return nil
}

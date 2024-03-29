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

func CompanyRoutes(db *sql.DB, r *mux.Router) error {
	companyRepository := repository.NewCompanyRepository(db)
	companyService := service.NewCompanyService(companyRepository)
	companyHandler := handler.NewCompanyHandler(companyService)

	r.HandleFunc("/company/register", companyHandler.RegisterCompany).Methods(http.MethodPost)
	r.HandleFunc("/company/info", middleware.Middleware(companyHandler.GetCompanyInfos)).Methods(http.MethodGet)

	return nil
}

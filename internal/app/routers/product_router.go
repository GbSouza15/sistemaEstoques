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

func ProductRoutes(db *sql.DB, r *mux.Router) error {
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	r.HandleFunc("/company/remove/product/{productId}", middleware.Middleware(productHandler.RemoveProduct)).Methods(http.MethodDelete)
	r.HandleFunc("/company/register/product/segment", middleware.Middleware(productHandler.AddProductSegment)).Methods(http.MethodPost)
	r.HandleFunc("/company/register/product", middleware.Middleware(productHandler.CreateProduct)).Methods(http.MethodPost)
	r.HandleFunc("/company/update/product", middleware.Middleware(productHandler.UpdateProduct)).Methods(http.MethodPut)
	r.HandleFunc("/company/search/product", middleware.Middleware(productHandler.SearchProduct)).Methods(http.MethodPost)

	return nil
}

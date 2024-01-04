package router

import (
	"database/sql"
	"fmt"
	"net/http"

	handler "github.com/GbSouza15/sistemaEstoque/internal/app/handlers"
	repository "github.com/GbSouza15/sistemaEstoque/internal/app/repositories"
	service "github.com/GbSouza15/sistemaEstoque/internal/app/services"
	"github.com/gorilla/mux"
)

func RoutesApi(db *sql.DB) error {
	r := mux.NewRouter()
	companyRepository := repository.NewCompanyRepository(db)
	companyService := service.NewCompanyService(companyRepository)
	companyHandler := handler.NewCompanyHandler(companyService)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	supplierRepository := repository.NewSupplierRepository(db)
	supplierService := service.NewSuppliersService(supplierRepository)
	supplierHandler := handler.NewSupplierHandler(supplierService)

	// Create supplier OK
	r.HandleFunc("/company/register/supplier", supplierHandler.CreateSupplier).Methods(http.MethodPost)

	// Get company infos Ok
	r.HandleFunc("/company/user", userHandler.GetCompanyInfos).Methods(http.MethodGet)

	// Delete product
	r.HandleFunc("/company/remove/product/{productId}", productHandler.RemoveProduct).Methods(http.MethodDelete)

	// Create product
	r.HandleFunc("/company/register/product", productHandler.CreateProduct).Methods(http.MethodPost)

	// Update product
	r.HandleFunc("/company/update/product", productHandler.UpdateProduct).Methods(http.MethodPut)

	// Register company OK
	r.HandleFunc("/company/register", companyHandler.RegisterCompany).Methods(http.MethodPost)

	// Register user OK
	r.HandleFunc("/company/register/user", userHandler.RegisterUser).Methods(http.MethodPost)

	// Login user OK
	r.HandleFunc("/company/user/login", userHandler.LoginUser).Methods(http.MethodPost)

	http.Handle("/", r)
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return err
	}

	return nil
}

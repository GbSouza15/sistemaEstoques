package router

import (
	"database/sql"
	"fmt"
	"net/http"

	handler "github.com/GbSouza15/sistemaEstoque/internal/app/handlers"
	repository "github.com/GbSouza15/sistemaEstoque/internal/app/repositories"
	service "github.com/GbSouza15/sistemaEstoque/internal/app/services"
	"github.com/GbSouza15/sistemaEstoque/pkg/middleware"
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

	depositRepository := repository.NewDepositRepository(db)
	depositService := service.NewDepositService(depositRepository)
	depositHandler := handler.NewDepositHandler(depositService)

	// Create deposit OK
	r.HandleFunc("/company/register/deposit", middleware.Middleware(depositHandler.CreateDeposit)).Methods(http.MethodPost)

	// Delete User
	r.HandleFunc("/company/remove/user/{userId}", middleware.Middleware(userHandler.DeleteUser)).Methods(http.MethodDelete)

	// Search product Ok
	r.HandleFunc("/company/search/product", middleware.Middleware(productHandler.SearchProduct)).Methods(http.MethodPost)

	// Create supplier OK
	r.HandleFunc("/company/register/supplier", middleware.Middleware(supplierHandler.CreateSupplier)).Methods(http.MethodPost)

	// Get company infos Ok
	r.HandleFunc("/company/user", middleware.Middleware(userHandler.GetCompanyInfos)).Methods(http.MethodGet)

	// Delete product Ok
	r.HandleFunc("/company/remove/product/{productId}", middleware.Middleware(productHandler.RemoveProduct)).Methods(http.MethodDelete)

	// Create product segment Ok
	r.HandleFunc("/company/register/product/segment", middleware.Middleware(productHandler.AddProductSegment)).Methods(http.MethodPost)

	// Create product OK
	r.HandleFunc("/company/register/product", middleware.Middleware(productHandler.CreateProduct)).Methods(http.MethodPost)

	// Update product OK
	r.HandleFunc("/company/update/product", middleware.Middleware(productHandler.UpdateProduct)).Methods(http.MethodPut)

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

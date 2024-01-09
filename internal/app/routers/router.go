package router

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RoutesApi(db *sql.DB) error {
	r := mux.NewRouter()

	if err := UserRoutes(db, r); err != nil {
		return err
	}

	if err := CompanyRoutes(db, r); err != nil {
		return err
	}

	if err := ProductRoutes(db, r); err != nil {
		return err
	}

	if err := DepositRoutes(db, r); err != nil {
		return err
	}

	if err := SupplierRoutes(db, r); err != nil {
		return err
	}

	http.Handle("/", r)
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return err
	}

	return nil
}

package main

import (
	"fmt"
	"os"

	router "github.com/GbSouza15/sistemaEstoque/internal/app/routers"
	"github.com/GbSouza15/sistemaEstoque/internal/database"
	"github.com/GbSouza15/sistemaEstoque/internal/database/schema"
)

func main() {

	db, err := database.InitDb()
	if err != nil {
		fmt.Printf("Error starting the database, %v ", err.Error())
		os.Exit(1)
	}

	err = schema.CreateSchema(db)
	if err != nil {
		fmt.Printf("Error in created schema %v ", err.Error())
	}

	err = router.RoutesApi(db)
	if err != nil {
		fmt.Printf("Error in router %v ", err.Error())
	}
}

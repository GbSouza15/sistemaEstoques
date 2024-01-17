package main

import (
	"fmt"
	"os"

	_ "github.com/GbSouza15/sistemaEstoque/cmd/docs"
	router "github.com/GbSouza15/sistemaEstoque/internal/app/routers"
	"github.com/GbSouza15/sistemaEstoque/internal/database"
	"github.com/GbSouza15/sistemaEstoque/internal/database/schema"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host
// @BasePath /
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

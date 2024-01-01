package schema

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func CreateSchema(db *sql.DB) error {
	godotenv.Load()
	schemaName := os.Getenv("DB_SCHEMA_NAME")

	schemaCreateString := fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schemaName)
	_, err := db.Exec(schemaCreateString)
	if err != nil {
		return err
	}

	fmt.Println("Schema created sucessfully")
	return nil
}

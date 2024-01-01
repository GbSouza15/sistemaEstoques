package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDb() (*sql.DB, error) {

	if err := godotenv.Load(); err != nil {
		fmt.Println(err.Error())
	}

	var (
		dbHost     = os.Getenv("DB_HOST")
		dbPort     = os.Getenv("DB_PORT")
		dbUser     = os.Getenv("DB_USER")
		dbPassword = os.Getenv("DB_PASSWORD")
		dbName     = os.Getenv("DB_NAME")
		dbSslmode  = os.Getenv("DB_SSLMODE")
	)

	connectString := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=%s", dbUser, dbName, dbPassword, dbHost, dbPort, dbSslmode)

	db, err := sql.Open("postgres", connectString)
	if err != nil {
		return nil, err
	}

	fmt.Println("Database connected")
	return db, nil
}

package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

type db struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
	db       *sql.DB
}

func ConnectDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var db db
	dbs, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {

		return nil, err
	}

	err = dbs.Ping()
	if err != nil {
		return nil, err
	}
	db.db = dbs
	return dbs, nil
}

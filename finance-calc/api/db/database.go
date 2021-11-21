package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"financeCalc/api/utils"
)

var DB *sql.DB

func Connect() {
	var err error

	log.Println("Starting database connection...")

	// open database
	DB, err = sql.Open("postgres", GetDatabaseConnectionString())
	utils.CheckError(err)

	// initial ping to ensure that we are all good!
	err = DB.Ping()
	utils.CheckError(err)

	log.Println("Database connected!")
}

func GetDbInstance() *sql.DB {
	return DB
}

func GetDatabaseConnectionString() string {
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)
}

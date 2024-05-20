package database

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var Db *sqlx.DB

func Connect() {
	var err error
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error while loading env")
		panic(err)
	}
	// godotenv.Load()
	Db, err = sqlx.Open("mysql", os.Getenv("SQL_PATH"))
	if err != nil {
		panic(err)
	}
}

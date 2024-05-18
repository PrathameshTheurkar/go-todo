package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func Connect() {
	var err error
	Db, err = sqlx.Open("mysql", "root:Prathamesh@1@tcp(localhost:3306)/go_todo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}

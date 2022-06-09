package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func OpenDBConnection() {
	var err error
	db, err = sql.Open("mysql", "maks:password@tcp(127.0.0.1:3306)/Film_Rec_System")

	if err != nil {
		log.Panicln(err.Error())
	}
}

func GetDBConnection() *sql.DB { return db }

type Search struct {
	Pattern string `binding:"max=100"`
	Genre   string `binding:"max=100"`
	Country string `binding:"max=100"`
}

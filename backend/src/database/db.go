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

func InsertStmt(query string, args []interface{}) error {
	var (
		insert *sql.Rows
		err    error
	)

	if args != nil {
		insert, err = db.Query(query)
	} else {
		insert, err = db.Query(query, args)
	}

	defer insert.Close()

	return err
}

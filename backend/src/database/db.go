package database

import (
	"../config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func OpenDBConnection(cnf *config.Config) {
	var err error
	db, err = sql.Open("mysql", cnf.UsernameDB+":"+cnf.PasswordDB+"@tcp("+cnf.AddressDB+")/"+cnf.NameDB)

	if err != nil {
		panic("Connection to db failed...")
	}
}

func GetDBConnection() *sql.DB { return db }

type Search struct {
	Pattern string `binding:"max=100"`
	Genre   string `binding:"max=100"`
	Country string `binding:"max=100"`
}

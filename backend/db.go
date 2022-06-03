package main

import (
	"crypto/rand"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
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

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}

	return salt, nil
}

func AddUser(user *User) error {
	salt, err := generateSalt()
	if err != nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword(
		append([]byte(user.Password), salt...),
		bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Salt = salt
	user.HashedPassword = hashedPassword

	err = InsertStmt("INSERT INTO USER(Username, Password) VALUES (?, ?)", []interface{}{user.Username, user.HashedPassword})
	if err != nil {
		return err
	}
	return err
}

func Authenticate(username, password string) (*User, error) {
	user := new(User)
	db := GetDBConnection()
	err := db.QueryRow("SELECT UserID, Password FROM User WHERE Username = ?", username).Scan(&user.ID, &user.HashedPassword)
	if err != nil {
		return nil, err
	}

	salted := append([]byte(password), user.Salt...)
	if err = bcrypt.CompareHashAndPassword(user.HashedPassword, salted); err != nil {
		return nil, err
	}
	return user, nil
}

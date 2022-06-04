package database

import (
	"crypto/rand"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	ID             int
	Username       string `binding:"required,min=5,max=63"`
	Password       string `binding:"required,min=7,max=63"`
	HashedPassword []byte `json:"-"`
	Salt           []byte `json:"-"`
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

	db := GetDBConnection()

	insert, err := db.Query("INSERT INTO USER(Username, Password, Hash) VALUES (?, ?, ?)", user.Username, user.HashedPassword, user.Salt)
	defer insert.Close()
	if err != nil {
		return err
	}
	return err
}

func Authenticate(username, password string) (*User, error) {
	user := new(User)
	db := GetDBConnection()
	err := db.QueryRow("SELECT UserID, Password, Hash FROM User WHERE Username = ?", username).Scan(&user.ID, &user.HashedPassword, &user.Salt)
	if err != nil {
		return nil, err
	}

	salted := append([]byte(password), user.Salt...)
	if err = bcrypt.CompareHashAndPassword(user.HashedPassword, salted); err != nil {
		return nil, err
	}
	return user, nil
}

func FetchUser(id int) (*User, error) {
	user := new(User)
	user.ID = id
	err := db.QueryRow("SELECT Username FROM User WHERE UserID = ?", id).Scan(&user.Username)
	if err != nil {
		log.Println("Error fetching user")
		return nil, err
	}
	return user, nil
}

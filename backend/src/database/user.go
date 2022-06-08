package database

import (
	"context"
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

	_, err = db.Exec("INSERT INTO User(Username, Password, Hash) VALUES (?, ?, ?)",
		user.Username, user.HashedPassword, user.Salt)
	if err != nil {
		return err
	}
	return err
}

func Authenticate(username, password string) (*User, error) {
	user := new(User)

	err := db.QueryRow("SELECT UserID, Password, Hash FROM User WHERE Username = ?", username).
		Scan(&user.ID, &user.HashedPassword, &user.Salt)
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
		log.Println("Error fetching user, i'm here")
		return nil, err
	}
	return user, nil
}

func DeleteUser(id int) error {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.ExecContext(ctx,
		"DELETE FROM Playlist AS p WHERE UserID = ? AND NOT EXISTS(SELECT * FROM Playlist_Film_INT AS `INT` WHERE p.PlaylistID = `INT`.PlaylistID)", id)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return errRollback
		}
		return err
	}

	_, err = tx.ExecContext(ctx,
		"DELETE FROM User WHERE UserID = ?", id)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return errRollback
		}
		return err
	}

	err = tx.Commit()
	return nil
}

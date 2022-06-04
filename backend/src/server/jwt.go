package server

import (
	"../database"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cristalhq/jwt"
	"log"
	"strconv"
	"time"
)

var (
	jwtSigner   jwt.Signer
	jwtVerifier jwt.Verifier
)

func jwtSetup() {
	var err error
	key := []byte("this is my secret")

	jwtSigner, err = jwt.NewSignerHS(jwt.HS256, key)
	if err != nil {
		log.Println("Error creating JWT signer : " + err.Error())
	}

	jwtVerifier, err = jwt.NewVerifierHS(jwt.HS256, key)
	if err != nil {
		log.Println("Error creating JWT verifier : " + err.Error())
	}
}

func generateJWT(user *database.User) string {
	claims := &jwt.RegisteredClaims{
		ID:        fmt.Sprint(user.ID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
	}
	builder := jwt.NewBuilder(jwtSigner)
	token, err := builder.Build(claims)
	if err != nil {
		log.Println("Error building JWT : " + err.Error())
	}
	return token.String()
}

func verifyJWT(tokenStr string) (int, error) {
	token, err := jwt.Parse([]byte(tokenStr), jwtVerifier)
	if err != nil {
		log.Println("Error parsing JWT verifier : " + tokenStr + " " + err.Error())
		return -1, err
	}

	if err := jwtVerifier.Verify(token); err != nil {
		log.Println("Error verifying token : " + err.Error())
		return -1, err
	}

	var claims jwt.RegisteredClaims
	if err := json.Unmarshal(token.Claims(), &claims); err != nil {
		log.Println("Error unmarshalling JWT claims : " + err.Error())
		return -1, err
	}

	if notExpired := claims.IsValidAt(time.Now()); !notExpired {
		return -1, errors.New("Token expired.")
	}

	id, err := strconv.Atoi(claims.ID)
	if err != nil {
		log.Println("Error converting claims ID to number : " + claims.ID + " " + err.Error())
		return -1, errors.New("ID in token is not valid")
	}
	return id, err
}

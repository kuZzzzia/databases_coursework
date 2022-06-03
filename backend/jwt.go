package main

import (
	"encoding/json"
	"errors"
	"fmt"
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

func generateJWT(user *User) string {
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
	token, err := jwt.Parse([]byte(tokenStr))
	if err != nil {
		log.Println("Error parsing JWT verifier : " + tokenStr + " " + err.Error())
		return 0, err
	}

	if err := jwtVerifier.Verify(token.Payload(), token.Signature()); err != nil {
		log.Println("Error verifying token : " + err.Error())
		return 0, err
	}

	var claims jwt.StandardClaims
	if err := json.Unmarshal(token.RawClaims(), &claims); err != nil {
		log.Println("Error unmarshalling JWT claims : " + err.Error())
		return 0, err
	}

	if notExpired := claims.IsValidAt(time.Now()); !notExpired {
		return 0, errors.New("Token expired.")
	}

	id, err := strconv.Atoi(claims.ID)
	if err != nil {
		log.Println("Error converting claims ID to number : " + claims.ID + " " + err.Error())
		return 0, errors.New("ID in token is not valid")
	}
	return id, err
}

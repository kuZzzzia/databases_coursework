package server

import (
	_ "crypto/rand"
	"film-network/src/database"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func signUp(ctx *gin.Context) {
	user := new(database.User)
	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "неверно заполнена форма"})
		return
	}
	if err := database.AddUser(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := database.Authenticate(user.Username, user.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "не удалось авторизоваться"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Signed up successfully.",
		"jwt": generateJWT(user),
	})
}

func signIn(ctx *gin.Context) {
	user := new(database.User)
	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "неверно заполнена форма"})
		return
	}
	user, err := database.Authenticate(user.Username, user.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "не удалось авторизоваться"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Signed in successfully.",
		"jwt": generateJWT(user),
	})
}

func authorization(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "пожалуйста, авторизуйтесь"})
		return
	}
	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "неудачная попытка авторизации"})
		return
	}
	if headerParts[0] != "Bearer" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "неудачная попытка авторизации"})
		return
	}
	userID, err := verifyJWT(headerParts[1])
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	user, err := database.FetchUser(userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "не удалось найти пользователя"})
		return
	}
	ctx.Set("user", user)
	ctx.Next()
}

func currentUser(ctx *gin.Context) (*database.User, error) {
	var err error
	_user, exists := ctx.Get("user")
	if !exists {
		log.Println("Current context user not set : " + err.Error())
		return nil, err
	}
	user, ok := _user.(*database.User)
	if !ok {
		log.Println("Current context user not set : " + err.Error())
		return nil, err
	}
	return user, nil
}

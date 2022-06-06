package server

import (
	"../database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func setUpRouter() *gin.Engine {
	// default router
	router := gin.Default()

	router.StaticFS("/images", http.Dir("./src/images"))

	api := router.Group("/api")
	{
		api.GET("")
		api.POST("/signUp", signUp)
		api.POST("/signIn", signIn)
		api.POST("/people", getPeople)
		api.POST("/films", getFilms)

	}

	main := router.Group("/")
	{
		main.GET("/person/:id", getPerson)
	}

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	return router
}

func Start() {
	jwtSetup()

	database.OpenDBConnection()
	defer database.GetDBConnection().Close()

	router := setUpRouter()

	router.Run(":8000")
}

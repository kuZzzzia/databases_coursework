package server

import (
	"../database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func setUpRouter() *gin.Engine {
	// default router
	router := gin.Default()

	//router.StaticFS("/images", http.Dir("../images"))

	api := router.Group("/api")
	{
		api.GET("")
		api.POST("/signUp", signUp)
		api.POST("/signIn", signIn)
	}

	main := api.Group("/")
	{

		main.POST("/people", getActors)
		main.POST("/films", getFilms)
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

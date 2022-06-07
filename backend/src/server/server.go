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

	auth := router.Group("/auth")
	auth.Use(authorization)
	{
		auth.POST("/film/:id", createMessage)
		auth.POST("/film/rate/:id", rateFilm)
		auth.POST("/playlist/rate/:id", ratePlaylist)
		auth.POST("/film/rateStatus/:id", getFilmRate)
		auth.POST("/playlist/rateStatus/:id", getPlaylistRate)
	}

	main := router.Group("/")
	{
		main.GET("/person/:id", getPerson)
		main.GET("/film/:id", getFilm)
		main.GET("/playlist/:id", getPlaylist)
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

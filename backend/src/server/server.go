package server

import (
	"../config"
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
		api.POST("/categories", getCategories)
	}

	auth := router.Group("/auth")
	auth.Use(authorization)
	{
		auth.POST("/film/:id", createMessage)
		auth.POST("/playlist/create", createPlaylist)
		auth.POST("/film/rate/:id", rate)
		auth.POST("/playlist/rate/:id", rate)
		auth.POST("/film/rateStatus/:id", getRating)
		auth.POST("/playlist/rateStatus/:id", getRating)
		auth.POST("/profile", getProfile)
		auth.POST("/delete/user", deleteUser)
		auth.POST("/delete/playlist/:id", deletePlaylist)
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

func Start(cnf *config.Config) {
	jwtSetup()

	database.OpenDBConnection(cnf)
	defer database.GetDBConnection().Close()

	router := setUpRouter()

	router.Run(cnf.PortServerBackend)
}

package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setUpRouter() *gin.Engine {
	// default router
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("")
		api.POST("/signUp", signUp)
		api.POST("/signIn", signIn)
	}

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	return router
}

func Start() {
	jwtSetup()

	OpenDBConnection()
	defer GetDBConnection().Close()

	router := setUpRouter()

	router.Run(":8000")
}

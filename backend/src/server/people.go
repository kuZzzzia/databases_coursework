package server

import (
	"../database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getActors(ctx *gin.Context) {
	search := new(database.Search)
	var err error
	if err = ctx.Bind(search); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var actors []*database.Person

	if actors, err = database.FetchPeople(search.Pattern); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": actors,
	})
}

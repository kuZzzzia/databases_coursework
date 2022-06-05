package server

import (
	"../database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getFilms(ctx *gin.Context) {
	search := new(database.Search)
	var err error
	if err = ctx.Bind(search); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var films []*database.Film

	if films, err = database.FetchFilms(search.Pattern); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": films,
	})
}

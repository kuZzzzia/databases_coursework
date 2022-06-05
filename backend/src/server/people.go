package server

import (
	"../database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getPeople(ctx *gin.Context) {
	search := new(database.Search)
	var err error
	if err = ctx.Bind(search); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var people []*database.Person

	if people, err = database.FetchPeople(search.Pattern); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if people == nil {
		people = []*database.Person{}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": people,
	})
}

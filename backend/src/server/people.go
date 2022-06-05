package server

import (
	"../database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func getPerson(ctx *gin.Context) {
	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Not valid ID."})
		return
	}
	person, roles, films, err := database.FetchPerson(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No person with such ID."})
		return
	}
	if roles == nil {
		roles = []*database.Role{}
	}
	if films == nil {
		films = []*database.Film{}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"person": person,
		"films":  films,
		"roles":  roles,
	})
}

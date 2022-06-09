package server

import (
	"film-network/src/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getPeople(ctx *gin.Context) {
	search := new(database.Search)
	var err error
	if err = ctx.Bind(search); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "неверный запрос"})
	}
	var people []*database.Person

	if people, err = database.FetchPeople(search.Pattern); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "не удалось найти людей"})
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "неверный запрос"})
		return
	}
	person, roles, films, err := database.FetchPerson(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "такой страницы не существует"})
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

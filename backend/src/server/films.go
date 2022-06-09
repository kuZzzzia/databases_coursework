package server

import (
	"../database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getFilms(ctx *gin.Context) {
	search := new(database.Search)
	var err error
	if err = ctx.Bind(search); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var films []*database.Film

	if films, err = database.FetchFilms(search); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if films == nil {
		films = []*database.Film{}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": films,
	})
}

func getFilm(ctx *gin.Context) {
	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Not valid ID."})
		return
	}

	film, cast, playlists, discussion, err := database.FetchFilm(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No person with such ID."})
		return
	}
	if cast == nil {
		cast = []*database.CastItem{}
	}
	if playlists == nil {
		playlists = []*database.Playlist{}
	}
	if discussion == nil {
		discussion = []*database.Message{}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"film":       film,
		"cast":       cast,
		"playlists":  playlists,
		"discussion": discussion,
	})
}

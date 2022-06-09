package server

import (
	"film-network/src/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getFilms(ctx *gin.Context) {
	search := new(database.Search)
	var err error
	if err = ctx.Bind(search); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "неверный запрос"})
	}
	var films []*database.Film

	if films, err = database.FetchFilms(search); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "не удалось найти фильмы"})
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "неверный запрос"})
		return
	}

	film, cast, playlists, discussion, err := database.FetchFilm(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "такого страницы не существует"})
		return
	}
	if cast == nil {
		cast = []*database.CastItem{}
	}
	if playlists == nil {
		playlists = []*database.Playlist{}
	}
	if discussion == nil {
		discussion = []*database.Comment{}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"film":       film,
		"cast":       cast,
		"playlists":  playlists,
		"discussion": discussion,
	})
}

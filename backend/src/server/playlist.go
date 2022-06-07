package server

import (
	"../database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getPlaylist(ctx *gin.Context) {
	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Not valid ID."})
		return
	}

	playlist, err := database.FetchPlaylist(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No person with such ID."})
		return
	}
	if playlist.Films == nil {
		playlist.Films = []*database.Film{}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"playlist": playlist,
	})
}

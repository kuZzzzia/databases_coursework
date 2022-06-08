package server

import (
	"../database"
	"github.com/gin-gonic/gin"
	"log"
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

func createPlaylist(ctx *gin.Context) {
	playlist := new(database.Playlist)
	if err := ctx.Bind(playlist); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	user, err := currentUser(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = database.AddPlaylist(playlist, user.ID)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Unable to create playlist"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Playlist created successfully.",
	})
}

func deletePlaylist(ctx *gin.Context) {
	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Not valid playlist ID."})
		return
	}

	user, err := currentUser(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = database.DeletePlaylist(id, user.ID)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Unable to delete playlist"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Playlist deleted successfully.",
	})
}

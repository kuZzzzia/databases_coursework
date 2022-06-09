package server

import (
	"film-network/src/database"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func getPlaylist(ctx *gin.Context) {
	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "неверный запрос"})
		return
	}

	playlist, err := database.FetchPlaylist(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "такой страницы не существует"})
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "неверный запрос"})
	}
	user, err := currentUser(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "нет доступа, авторизуйтесь"})
		return
	}
	err = database.AddPlaylist(playlist, user.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "не удалось создать подборку"})
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "неверный запрос"})
		return
	}

	user, err := currentUser(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "нет доступа, авторизуйтесь"})
		return
	}
	err = database.DeletePlaylist(id, user.ID)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "не удалось удалить подборку"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Playlist deleted successfully.",
	})
}

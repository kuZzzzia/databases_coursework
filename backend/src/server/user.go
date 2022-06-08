package server

import (
	"../database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getProfile(ctx *gin.Context) {
	user, err := currentUser(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	playlists, _ := database.FetchPlaylists(database.PlaylistsForProfile, user.ID)
	if playlists == nil {
		playlists = []*database.Playlist{}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":       "Profile fetched successfully.",
		"username":  user.Username,
		"playlists": playlists,
	})
}

func deleteUser(ctx *gin.Context) {
	user, err := currentUser(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = database.DeleteUser(user.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "User deleted successfully.",
	})
}

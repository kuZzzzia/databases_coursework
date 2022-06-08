package server

import (
	"../database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func rate(ctx *gin.Context) {
	rate := new(database.Rate)
	if err := ctx.Bind(rate); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Not valid" + rate.Src + "ID."})
		return
	}

	user, err := currentUser(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var query string
	if rate.Src == "film" {
		query = database.AddRatingToFilm
	} else if rate.Src == "playlist" {
		query = database.AddRatingToPlaylist
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "This service is unavailable"})
		return
	}
	if err := database.AddRate(query, user.ID, rate.Like, id); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": rate.Src + " rated successfully.",
	})
}

func getRating(ctx *gin.Context) {
	rate := new(database.Rate)
	if err := ctx.Bind(rate); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Not valid" + rate.Src + "ID."})
		return
	}
	user, err := currentUser(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var query string
	if rate.Src == "film" {
		query = database.GetUserRatingOfFilm
	} else if rate.Src == "playlist" {
		query = database.GetUserRatingOfPlaylist
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "This service is unavailable"})
		return
	}
	rating, err := database.GetUserRate(query, user.ID, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  rate.Src + " rate queried successfully.",
		"rate": rating,
	})
}

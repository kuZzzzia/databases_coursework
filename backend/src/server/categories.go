package server

import (
	"../database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getCategories(ctx *gin.Context) {
	genres, countries, err := database.FetchCategories()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if genres == nil {
		genres = []*string{}
	}
	if countries == nil {
		countries = []*string{}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"countries": countries,
		"genres":    genres,
	})
}

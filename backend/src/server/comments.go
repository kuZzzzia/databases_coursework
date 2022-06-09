package server

import (
	"film-network/src/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func createMessage(ctx *gin.Context) {
	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "неверный запрос"})
		return
	}
	post := new(database.Comment)
	if err := ctx.Bind(post); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "неверно заполнена форма"})
		return
	}
	user, err := currentUser(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "нет доступа, авторизуйтесь"})
		return
	}
	if err := database.AddComment(user, post, id); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "не удалось добавить комментарий"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Post created successfully.",
		"post": post,
	})
}

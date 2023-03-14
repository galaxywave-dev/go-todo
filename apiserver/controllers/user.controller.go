package controllers

import (
	"net/http"

	"galaxywave.com/go-todo/apiserver/models"
	"github.com/gin-gonic/gin"
)

func GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": models.FilteredResponse(&currentUser)}})
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

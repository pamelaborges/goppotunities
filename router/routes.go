package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initialize(router *gin.Engine) {

	v1 := router.Group("/api/v1")

	v1.GET("/opening", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	v1.POST("/opening", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	v1.PUT("/opening", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	v1.DELETE("/opening", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

}

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pamelaborges/goppotunities/handler"
)

func initialize(router *gin.Engine) {
	path := "/opening"
	v1 := router.Group("/api/v1")
	{
		v1.GET(path, handler.GetOpeningHandler)
		v1.POST(path, handler.CreateOpeningHandler)
		v1.PUT(path, handler.UpdateOpeningHandler)
		v1.DELETE(path, handler.DeleteOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)
	}
}

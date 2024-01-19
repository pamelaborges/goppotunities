package router

import "github.com/gin-gonic/gin"

func Init() {
	router := gin.Default()
	initialize(router)
	router.Run(":8080")
}

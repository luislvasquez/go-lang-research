package api

import (
	"github.com/gin-gonic/gin"
	networks "main/api/networks"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	setUpRouter(router)
	return router
}

func setUpRouter(router *gin.Engine) {
	api := router.Group("/api")
	{
		networks.RegisterRouter(api.Group("/networks"))
	}
}
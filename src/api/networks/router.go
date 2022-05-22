package networks

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup) {
	r.POST("/", addNetworks)
	r.GET("/", getNetworks)
}
package networks

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup) {
	r.POST("/", addNetworks)
	r.POST("/async", addNetworksAsync)
	r.GET("/", getNetworks)
}
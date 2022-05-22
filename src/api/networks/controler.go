package networks

import (
	"github.com/gin-gonic/gin"
)

func addNetworks(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "Hello from addNetworks"})
}

func getNetworks(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "Hello from getNetworks"})
}

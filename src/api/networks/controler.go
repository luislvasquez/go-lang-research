package networks

import (
	"time"
	"github.com/gin-gonic/gin"
)

func addNetworks(c *gin.Context) {
	start := time.Now()
	
	create(c)

	elapsedTime := float64(time.Since(start).Milliseconds())/float64(1000)

	c.JSON(200, gin.H{"msg": "Network created", "responseTimeInSeconds": elapsedTime})
}

func addNetworksAsync(c *gin.Context) {
	start := time.Now()

	createAsync(c)

	elapsedTime := float64(time.Since(start).Milliseconds())/float64(1000)
	c.JSON(200, gin.H{"msg": "Network creation triggered", "responseTimeInSeconds": elapsedTime})
}

func getNetworks(c *gin.Context) {
    c.IndentedJSON(200, readAll())
}

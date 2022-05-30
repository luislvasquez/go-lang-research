package networks

import (
	"time"
	"github.com/gin-gonic/gin"
)

func addNetworks(c *gin.Context) {
	start := time.Now()
	
	response_status, response_msg := create(c)

	elapsedTime := float64(time.Since(start).Milliseconds())/float64(1000)

	c.JSON(response_status, gin.H{"msg": response_msg, "responseTimeInSeconds": elapsedTime})
}

func addNetworksAsync(c *gin.Context) {
	start := time.Now()
	
	response_status, response_msg := createAsync(c)

	elapsedTime := float64(time.Since(start).Milliseconds())/float64(1000)

	c.JSON(response_status, gin.H{"msg": response_msg, "responseTimeInSeconds": elapsedTime})
}

func getNetworks(c *gin.Context) {
    c.IndentedJSON(200, readAll())
}

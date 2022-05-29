package networks

import (
	"time"
	"github.com/gin-gonic/gin"
)
 const DB_WRITTING_TIME_IN_SECONDS = 10

type Network struct {
	Name string
	Status string
}

var networksTable = []Network{
	{Name: "network_01", Status: "ONLINE"},
	{Name: "network_02", Status: "ONLINE"},
}

func create(c *gin.Context) {
	var newNetwork Network
	c.BindJSON(&newNetwork)
	if newNetwork.Name == "" {
		return
	}

	/* CPU demanding CRUD operation starts */
	time.Sleep(DB_WRITTING_TIME_IN_SECONDS * time.Second)
	newNetwork.Status = "ONLINE"
	networksTable = append(networksTable, newNetwork)
}

func createAsync(c *gin.Context) {
	var newNetwork Network
	c.BindJSON(&newNetwork)
	if newNetwork.Name == "" {
		return
	}
	newNetwork.Status = "PREPARING"
	networksTable = append(networksTable, newNetwork)

	/* CPU demanding CRUD operation starts, but now asyc */
	go func() {
		time.Sleep(DB_WRITTING_TIME_IN_SECONDS * time.Second)
		for i, network := range networksTable {
			if network.Name == newNetwork.Name {
				networksTable = append(networksTable[:i], networksTable[i+1:]...)  // Deletes placeholder network
				// Creating the real one as "ONLINE"
				var modifNetwork Network
				modifNetwork.Name = newNetwork.Name
				modifNetwork.Status = "ONLINE"
				networksTable = append(networksTable, modifNetwork)
				return
			}
		}
	} ()
}

func readAll() []Network{
	/*
	Technically not needed since it would only read
	from the list, but this could be a DB execution
	*/
	return networksTable
}
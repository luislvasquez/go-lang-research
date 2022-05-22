package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type album struct {
	ID	string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float32 `json:"price"`
}

var albums = []album {
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
    {ID: "4", Title: "Lord of the Rings", Artist: "J.R.R. Tolkien", Price: 59.99},
}

func main() {
	r := gin.New()
	r.GET("/albums", func(c *gin.Context) {
		c.JSON(http.StatusOK, albums)
	})
	r.Run()
}

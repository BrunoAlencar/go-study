package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 5600},
	{ID: "2", Title: "A Love Supreme", Artist: "John Coltrane", Price: 7300},
	{ID: "3", Title: "Kind of Blue", Artist: "Miles Davis", Price: 6300},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.DELETE("/albums/:id", deleteAlbumById)
	router.PUT("/albums/:id", putAlbumById)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		log.Println(err)
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func deleteAlbumById(c *gin.Context) {
	id := c.Param("id")

	for i, album := range albums {
		if album.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func putAlbumById(c *gin.Context) {
	id := c.Param("id")
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		log.Println(err)
		return
	}

	for i, album := range albums {
		if album.ID == id {
			albums[i] = newAlbum
			c.IndentedJSON(http.StatusOK, newAlbum)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

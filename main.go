package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Album represents data about a record album.
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Create a new struct to handle API responses
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty"`
}

// albums slice to seed record album data.
var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// GetAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	response := APIResponse{
		Success: true,
		Data:    albums,
	}
	c.JSON(http.StatusOK, response)
}

// GetAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			response := APIResponse{
				Success: true,
				Data:    a,
			}
			c.JSON(http.StatusOK, response)
			return
		}
	}

	response := APIResponse{
		Success: false,
		Error:   "Album not found",
	}
	c.JSON(http.StatusNotFound, response)
}

// PostAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var newAlbum Album

	// Call ShouldBindJSON to bind the received JSON to newAlbum.
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		response := APIResponse{
			Success: false,
			Error:   err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	response := APIResponse{
		Success: true,
		Data:    newAlbum,
	}
	c.JSON(http.StatusCreated, response)
}

// UpdateAlbum updates an existing album with new data.
func UpdateAlbum(c *gin.Context) {
	id := c.Param("id")

	var updatedAlbum Album
	if err := c.ShouldBindJSON(&updatedAlbum); err != nil {
		response := APIResponse{
			Success: false,
			Error:   err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Find the album in the slice and update it.
	for i, a := range albums {
		if a.ID == id {
			albums[i] = updatedAlbum
			response := APIResponse{
				Success: true,
				Data:    updatedAlbum,
			}
			c.JSON(http.StatusOK, response)
			return
		}
	}

	response := APIResponse{
		Success: false,
		Error:   "Album not found",
	}
	c.JSON(http.StatusNotFound, response)
}

// DeleteAlbum deletes an album with the given ID.
func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")

	// Find the album in the slice and remove it.
	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			response := APIResponse{
				Success: true,
				Data:    nil,
			}
			c.JSON(http.StatusOK, response)
			return
		}
	}

	response := APIResponse{
		Success: false,
		Error:   "Album not found",
	}
	c.JSON(http.StatusNotFound, response)
}

func main() {
	router := gin.Default()
	router.GET("/albums", GetAlbums)
	router.GET("/albums/:id", GetAlbumByID)
	router.POST("/albums", PostAlbums)
	router.PUT("/albums/:id", UpdateAlbum)
	router.DELETE("/albums/:id", DeleteAlbum)

	router.Run("localhost:8085")
}

package main

import (
	"net/http"

  "github.com/joaaomanooel/booksgo-api/models"
  "github.com/joaaomanooel/booksgo-api/controllers"

  "github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

  // Routes
	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"data": "Wellcome to Books Go API"})
  })

	// Run the server
	r.Run()
}

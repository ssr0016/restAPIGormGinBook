package main

import (
	"rest1/controllers"
	"rest1/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)         // new
	r.POST("/books", controllers.CreateBook)       // new
	r.GET("/books/:id", controllers.FindBook)      // new
	r.PATCH("/books/:id", controllers.UpdateBook)  // new
	r.DELETE("/books/:id", controllers.DeleteBook) // new

	err := r.Run()
	if err != nil {
		return
	}
}

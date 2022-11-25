package main

import (
	"gin-test/controllers"
	"gin-test/middlewares"
	"gin-test/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	protected := r.Group("/api/auth")
	protected.Use(middlewares.JwtAuthMiddleware())

	protected.GET("/books", controllers.FindBooks) // new
	protected.POST("/books", controllers.CreateBook)
	protected.GET("/books/:id", controllers.FindBook)
	protected.PUT("/books/:id", controllers.UpdateBook)

	protected.GET("/users", controllers.FindUsers)
	protected.GET("/users/:id", controllers.FindUser)

	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}

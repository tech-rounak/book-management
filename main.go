package main

import (
	"github/tech-rounak/book-management/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.BookRoutes(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Server is running"})
	})
	router.Run(":" + port)
}

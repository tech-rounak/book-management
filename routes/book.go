package routes

import (
	"github/tech-rounak/book-management/controllers"

	"github.com/gin-gonic/gin"
)

func BookRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/books/create", controllers.CreateBook())
	incomingRoutes.PUT("/books/update/:bookId", controllers.UpdateBook())
	incomingRoutes.GET("/book/:bookId", controllers.GetBookById())
	incomingRoutes.DELETE("/books/delete/:bookId", controllers.DeleteBook())
}

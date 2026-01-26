package router

import (
	"book-shop/internal/author"
	"book-shop/internal/book"

	"github.com/gin-gonic/gin"
)

func SetupRouter(authorHandler *author.Handler, bookHandler *book.Handler) *gin.Engine {
	router := gin.Default()

	// Ping endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Author routes
	authorGroup := router.Group("/authors")
	{
		authorGroup.POST("", authorHandler.CreateAuthor)
		authorGroup.GET("", authorHandler.GetAllAuthors)
		authorGroup.GET("/:id", authorHandler.GetAuthorByID)
		authorGroup.PUT("/:id", authorHandler.UpdateAuthor)
		authorGroup.DELETE("/:id", authorHandler.DeleteAuthor)
		authorGroup.GET("/:id/books", authorHandler.GetAuthorByIDWithBooks)
	}

	// Book routes
	bookGroup := router.Group("/books")
	{
		bookGroup.POST("", bookHandler.CreateBook)
		bookGroup.GET("", bookHandler.GetAllBooks)
		bookGroup.GET("/:id", bookHandler.GetBookByID)
		bookGroup.PUT("/:id", bookHandler.UpdateBook)
		bookGroup.DELETE("/:id", bookHandler.DeleteBook)
		bookGroup.GET("/:id/authors", bookHandler.GetBookByIDWithAuthors)
	}

	return router
}

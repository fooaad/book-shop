package main

import (
	"book-shop/internal/author"
	"book-shop/internal/book"
	"book-shop/internal/database"
	"book-shop/internal/router"
	"log"
)

func main() {
	// Initialize database
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize repositories
	authorRepo := author.NewRepository(db)
	bookRepo := book.NewRepository(db)

	// Initialize services
	authorService := author.NewService(authorRepo)
	bookService := book.NewService(bookRepo)

	// Initialize handlers
	authorHandler := author.NewHandler(authorService)
	bookHandler := book.NewHandler(bookService)

	// Set up router
	r := router.SetupRouter(authorHandler, bookHandler)

	r.Run() // listens on 0.0.0.0:8080 by default
}

package bookstrap

import (
	"book-store/internal/application/command"
	"book-store/internal/application/query"
	"book-store/internal/domain/book"
	"book-store/internal/infrastructure/http/handler"
	"book-store/internal/infrastructure/postgres"
	"log"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// TODO: Move to configuration
	connStr := "postgresql://jvm:admin@localhost:5432/jvm"

	// Initialize database connection
	dbConn, err := postgres.InitDB(connStr)
	if err != nil {
		log.Fatal("failed to initialize database:", err)
	}

	bookRepo := postgres.NewBookRepository(dbConn.GetDB())
	bookService := book.NewService(bookRepo)

	createBookHandler := command.NewCreateBookHandler(bookService)
	getBookByISBNHandler := query.NewGetBookByISBNHandler(bookService)

	bookHandler := handler.NewBookHandler(createBookHandler, getBookByISBNHandler)

	router.POST("/books", bookHandler.CreateBook)
	router.GET("/books/:isbn", bookHandler.GetBookByISBN)
	// router.GET("/books", bookHandler.ListBooks)
	// router.PUT("/books/:id", bookHandler.UpdateBook)
	// router.DELETE("/books/:id", bookHandler.DeleteBook)

	return router
}

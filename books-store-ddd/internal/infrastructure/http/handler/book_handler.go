package handler

import (
	"book-store/internal/application/command"
	"book-store/internal/application/query"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	createBookHandler    *command.CreateBookHandler
	getBookByISBNHandler *query.GetBookByISBNHandler
}

func NewBookHandler(createBookHandler *command.CreateBookHandler, getBookByISBNHandler *query.GetBookByISBNHandler) *BookHandler {
	return &BookHandler{
		createBookHandler:    createBookHandler,
		getBookByISBNHandler: getBookByISBNHandler,
	}
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var req command.CreateBook
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.createBookHandler.Handle(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully"})
}

func (h *BookHandler) GetBookByISBN(c *gin.Context) {
	isbn := c.Param("isbn")
	query := query.GetBookByISBN{ISBN: isbn}

	book, err := h.getBookByISBNHandler.Handle(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

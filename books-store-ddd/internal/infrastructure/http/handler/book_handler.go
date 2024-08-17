package handler

import (
	"book-store/internal/application/command"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	createBookHandler *command.CreateBookHandler
}

func NewBookHandler(createBookHandler *command.CreateBookHandler) *BookHandler {
	return &BookHandler{
		createBookHandler: createBookHandler,
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

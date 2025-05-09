package controllers

import (
	"net/http"

	"go-learn/gin-crud/models"

	"github.com/gin-gonic/gin"
)

// BookController handles book-related operations
type BookController struct {
	DB *models.BookDB
}

// NewBookController creates a new book controller
func NewBookController() *BookController {
	return &BookController{
		DB: models.NewBookDB(),
	}
}

// GetBooks handles GET /books
func (c *BookController) GetBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.DB.GetAll())
}

// GetBook handles GET /books/:id
func (c *BookController) GetBook(ctx *gin.Context) {
	id := ctx.Param("id")
	book, err := c.DB.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, book)
}

// CreateBook handles POST /books
func (c *BookController) CreateBook(ctx *gin.Context) {
	var book models.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.DB.Create(&book)
	ctx.JSON(http.StatusCreated, book)
}

// UpdateBook handles PUT /books/:id
func (c *BookController) UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	var book models.Book
	
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	book.ID = id
	if err := c.DB.Update(id, &book); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, book)
}

// DeleteBook handles DELETE /books/:id
func (c *BookController) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	
	if err := c.DB.Delete(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
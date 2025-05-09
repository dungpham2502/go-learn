package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	db "go-learn/gin-crud/db/sqlc"
	"go-learn/gin-crud/models"
	"go-learn/gin-crud/db/store"
)

// BookController handles requests related to book_store
type BookController struct {
	store *store.Store
}

func NewBookController(store *store.Store) *BookController {
	return &BookController{
		store: store,
	}
}

// GetBooks handles GET /books
// This endpoint lists all books with pagination
func (c *BookController) GetBooks(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.Query("limit", 10))
	offset, _ := strconv.Atoi(ctx.Query("offset", 0))

	// Validate pagination parameters
	if limit < 1 || limit > 100 {
		limit = 10 // Default limit
	}
	if offset < 0 {
		offset = 0 // Default offset
	}
	// Query databse using SQLC-generated code
	books, err := c.store.ListBooks(ctx, db.ListBooksParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}

	// Convert to response DTOs
	response := make([]models.BookResponse, len(books))
	for i, book := range books {
		response[i] = models.BookResponse{book}
	}
	ctx.JSON(http.StatusOK, response)
}

// GetBook handles GET /books/:id


// CreateBook handles POST /books
// This endpoint creates a new book
func (c *BookController) CreateBook(ctx *gin.Context) {
	var req models.BookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if ISBN already exists
	_, err := c.store.GetBookByISBN(ctx, req.ISBN)
	if err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "ISBN already exists"})
		return
	} else if err != sql.ErrNoRows {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check ISBN"})
		return
	}
	// Create book using SQLC-generated code
	book, err := c.store.CreateBook(ctx, db.CreateBookParams{
		Title:         req.Title,
		Author: 	  req.Author,
		ISBN:          req.ISBN,
		PublishedYear: req.PublishedYear,
		Price:         req.Price,
		Quantity:      req.Quantity,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}
	// Return the created book
	ctx.JSON(http.StatusCreated, models.BookResponse(book))
	
}
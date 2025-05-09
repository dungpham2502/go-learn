package models

import (
	"time"
)

type BookRequest struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	ISBN   string `json:"isbn" binding:"required"`
	PublishedYear int32  `json:"published_year" binding:"required"`
	Price  string `json:"price" binding:"required"`
	Quantity int32  `json:"quantity" binding:"required"`
}

// BookResponse represents the book data sent back to clients
type BookResponse struct {
    ID            int32     `json:"id"`
    Title         string    `json:"title"`
    Author        string    `json:"author"`
    ISBN          string    `json:"isbn"`
    PublishedYear int32     `json:"published_year"`
    Price         string    `json:"price"`
    Quantity      int32     `json:"quantity"`
    CreatedAt     time.Time `json:"created_at"`
}

// UserRequest represents data needed to create/update a user
type UserRequest struct {
    Name     string `json:"name" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

// UserResponse represents the user data sent back to clients
// Note: password is omitted for security
type UserResponse struct {
    ID        int32     `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
}


// Convert DB models to response DTOs
func NewBookResponse(book db.Book) BookResponse {
    return BookResponse{
        ID:            book.ID,
        Title:         book.Title,
        Author:        book.Author,
        ISBN:          book.Isbn,
        PublishedYear: book.PublishedYear,
        Price:         book.Price,
        Quantity:      book.Quantity,
        CreatedAt:     book.CreatedAt,
    }
}

func NewUserResponse(user db.User) UserResponse {
    return UserResponse{
        ID:        user.ID,
        Name:      user.Name,
        Email:     user.Email,
        CreatedAt: user.CreatedAt,
    }
}


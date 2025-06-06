// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package db

import (
	"database/sql"
	"time"
)

type Book struct {
	ID            int32     `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	Isbn          string    `json:"isbn"`
	PublishedYear int32     `json:"published_year"`
	Price         string    `json:"price"`
	Quantity      int32     `json:"quantity"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Loan struct {
	ID         int32        `json:"id"`
	UserID     int32        `json:"user_id"`
	BookID     int32        `json:"book_id"`
	BorrowedAt time.Time    `json:"borrowed_at"`
	DueDate    time.Time    `json:"due_date"`
	ReturnedAt sql.NullTime `json:"returned_at"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
}

type User struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

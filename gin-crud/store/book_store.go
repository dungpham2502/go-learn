package store

import (
	"context"
	"errors"
	"fmt"
	"time"

	db "go-learn/gin-crud/db/sqlc"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// BorrowBookParams contains the parameters for borrowing a book
type BorrowBookParams struct {
	UserID int32 `json:"user_id"`
	BookID int32 `json:"book_id"`
	DueDate time.Time `json:"due_date"`
}


// BorrowBookResult contains the result of the borrow transaction
type BorrowBookResult struct {
	Book db.Book `json:"book"`
	Loan db.Loan `json:"loan"`
}

// BorrowBookTX handles the borrowing of a book transaction
func (store *Store) BorrowBookTX(ctx context.Context, arg BorrowBookParams) (BorrowBookResult, error) {
	var result BorrowBookResult
	err := store.execTx(ctx, func(q *db.Queries) error {
		var err error
		
		// 1. Get book to check availability
		book, err := q.GetBook(ctx, arg.BookID)
		if err != nil {
			return fmt.Errorf("failed to get book: %w", err)
		}

		// 2. Check if the book is available
		if book.Quantity <= 0 {
			return errors.New("book is not available")
		}

		// 3. Reduce the book quantity
		updateBookParams := db.UpdateBookParams{
			ID: book.ID,
			Title: book.Title,
			Author: book.Author,
			Price: book.Price,
			Quantity: book.Quantity - 1,
		}
		updateBook, err := q.UpdateBook(ctx, updateBookParams)
		if err != nil {
			return fmt.Errorf("failed to update book: %w", err)
		}

		// 4. Create a loan record
		createLoanParams := db.CreateLoanParams{
			UserID: arg.UserID,
			BookID: arg.BookID,
			DueDate: arg.DueDate,
		}
		loan, err := q.CreateLoan(ctx, createLoanParams)
		if err != nil {
			return fmt.Errorf("failed to create loan: %w", err)
		}
		// 5. Set the result
		result.Book = updateBook
		result.Loan = loan
		return nil
	})

	return result, err
}
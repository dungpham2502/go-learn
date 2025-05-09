package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	db "go-learn/gin-crud/db/sqlc"
	_ "github.com/lib/pq"
)

// Store provides all functions to execute database queries and transactions
type Store struct {
	*db.Queries // Embed the queries struct from sqlc
	db *sql.DB // Keep a reference to the database connection
}

// New connection establishes a new database connection
func NewConnection(dbSource string) (*sql.DB, error){
	conn, err := sql.Open("postgres", dbSource)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to db: %w", err)
	}

	// Set connection pool parameters
	conn.SetMaxIdleConns(10) //Maximum number of idle connections
	conn.SetMaxOpenConns(100) //Maximum number of open connections
	conn.SetConnMaxLifetime(time.Hour) //Maximum lifetime of a connection

	// Verify the connection
	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("cannot ping db: %w", err)
	}

	return conn, nil
}

// NewStore creates a new Store instance
func NewStore(dbConn *sql.DB) *Store {
	return &Store{
		Queries: db.New(dbConn), // Initialize the Queries struct
		db:      dbConn,        // Store the database connection
	}
}

// execTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*db.Queries) error) error {
    tx, err := store.db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }

    q := db.New(tx)
    err = fn(q)
    if err != nil {
        // Attempt rollback, but return original error
        if rbErr := tx.Rollback(); rbErr != nil {
            return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
        }
        return err
    }

    return tx.Commit()
}



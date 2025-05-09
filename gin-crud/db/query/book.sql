-- name: CreateBook :one
INSERT INTO books (
  title,
  author,
  isbn,
  published_year,
  price,
  quantity
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetBook :one
SELECT * FROM books
WHERE id = $1 LIMIT 1;

-- name: GetBookByISBN :one
SELECT * FROM books
WHERE isbn = $1 LIMIT 1;

-- name: ListBooks :many
SELECT * FROM books
ORDER BY title
LIMIT $1
OFFSET $2;

-- name: SearchBooks :many
SELECT * FROM books
WHERE 
  title ILIKE '%' || $1 || '%' OR
  author ILIKE '%' || $1 || '%'
ORDER BY title
LIMIT $2
OFFSET $3;

-- name: UpdateBook :one
UPDATE books
SET 
  title = $2,
  author = $3,
  price = $4,
  quantity = $5,
  updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;
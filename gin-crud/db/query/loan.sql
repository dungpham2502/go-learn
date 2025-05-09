-- name: CreateLoan :one
INSERT INTO loans (
  user_id,
  book_id,
  due_date
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetLoan :one
SELECT * FROM loans
WHERE id = $1 LIMIT 1;

-- name: ListLoansByUser :many
SELECT * FROM loans
WHERE user_id = $1
ORDER BY borrowed_at DESC
LIMIT $2
OFFSET $3;

-- name: ListActiveLoans :many
SELECT * FROM loans
WHERE returned_at IS NULL
ORDER BY due_date ASC
LIMIT $1
OFFSET $2;

-- name: ListOverdueLoans :many
SELECT * FROM loans
WHERE 
  returned_at IS NULL AND
  due_date < now()
ORDER BY due_date ASC
LIMIT $1
OFFSET $2;

-- name: ReturnBook :one
UPDATE loans
SET 
  returned_at = now(),
  updated_at = now()
WHERE 
  id = $1 AND
  returned_at IS NULL
RETURNING *;

-- name: DeleteLoan :exec
DELETE FROM loans
WHERE id = $1;
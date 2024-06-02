-- name: CreateTransfer :one
INSERT INTO transfers (
  form_account_id,
  to_account_id,
  amount
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: ListTransfersByFormAccount :many
SELECT * FROM transfers
WHERE form_account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: ListTransfersByToAccount :many
SELECT * FROM transfers
WHERE to_account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateTransfer :one
UPDATE transfers
SET form_account_id = $2,
    to_account_id = $3,
    amount = $4
WHERE id = $1
RETURNING *;

-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = $1;
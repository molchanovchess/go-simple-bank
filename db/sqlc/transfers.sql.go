// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: transfers.sql

package db

import (
	"context"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO transfers (
  form_account_id,
  to_account_id,
  amount
) VALUES (
  $1, $2, $3
) RETURNING id, form_account_id, to_account_id, amount, created_at
`

type CreateTransferParams struct {
	FormAccountID int64 `json:"form_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, createTransfer, arg.FormAccountID, arg.ToAccountID, arg.Amount)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FormAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTransfer = `-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = $1
`

func (q *Queries) DeleteTransfer(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTransfer, id)
	return err
}

const getTransfer = `-- name: GetTransfer :one
SELECT id, form_account_id, to_account_id, amount, created_at FROM transfers
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTransfer(ctx context.Context, id int64) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTransfer, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FormAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listTransfers = `-- name: ListTransfers :many
SELECT id, form_account_id, to_account_id, amount, created_at FROM transfers
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListTransfersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, listTransfers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transfer
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FormAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTransfersByFormAccount = `-- name: ListTransfersByFormAccount :many
SELECT id, form_account_id, to_account_id, amount, created_at FROM transfers
WHERE form_account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListTransfersByFormAccountParams struct {
	FormAccountID int64 `json:"form_account_id"`
	Limit         int32 `json:"limit"`
	Offset        int32 `json:"offset"`
}

func (q *Queries) ListTransfersByFormAccount(ctx context.Context, arg ListTransfersByFormAccountParams) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, listTransfersByFormAccount, arg.FormAccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transfer
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FormAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTransfersByToAccount = `-- name: ListTransfersByToAccount :many
SELECT id, form_account_id, to_account_id, amount, created_at FROM transfers
WHERE to_account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListTransfersByToAccountParams struct {
	ToAccountID int64 `json:"to_account_id"`
	Limit       int32 `json:"limit"`
	Offset      int32 `json:"offset"`
}

func (q *Queries) ListTransfersByToAccount(ctx context.Context, arg ListTransfersByToAccountParams) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, listTransfersByToAccount, arg.ToAccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transfer
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FormAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTransfer = `-- name: UpdateTransfer :one
UPDATE transfers
SET form_account_id = $2,
    to_account_id = $3,
    amount = $4
WHERE id = $1
RETURNING id, form_account_id, to_account_id, amount, created_at
`

type UpdateTransferParams struct {
	ID            int64 `json:"id"`
	FormAccountID int64 `json:"form_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

func (q *Queries) UpdateTransfer(ctx context.Context, arg UpdateTransferParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, updateTransfer,
		arg.ID,
		arg.FormAccountID,
		arg.ToAccountID,
		arg.Amount,
	)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FormAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}
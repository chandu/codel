// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: flows.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createFlow = `-- name: CreateFlow :one
INSERT INTO flows (
  name, status, container_id
)
VALUES (
  $1, $2, $3
)
RETURNING id, created_at, updated_at, name, status, container_id
`

type CreateFlowParams struct {
	Name        pgtype.Text
	Status      pgtype.Text
	ContainerID pgtype.Int8
}

func (q *Queries) CreateFlow(ctx context.Context, arg CreateFlowParams) (Flow, error) {
	row := q.db.QueryRow(ctx, createFlow, arg.Name, arg.Status, arg.ContainerID)
	var i Flow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Status,
		&i.ContainerID,
	)
	return i, err
}

const readAllFlows = `-- name: ReadAllFlows :many
SELECT
  f.id, f.created_at, f.updated_at, f.name, f.status, f.container_id,
  c.name AS container_name
FROM flows f
LEFT JOIN containers c ON f.container_id = c.id
ORDER BY f.created_at DESC
`

type ReadAllFlowsRow struct {
	ID            int64
	CreatedAt     pgtype.Timestamp
	UpdatedAt     pgtype.Timestamp
	Name          pgtype.Text
	Status        pgtype.Text
	ContainerID   pgtype.Int8
	ContainerName pgtype.Text
}

func (q *Queries) ReadAllFlows(ctx context.Context) ([]ReadAllFlowsRow, error) {
	rows, err := q.db.Query(ctx, readAllFlows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ReadAllFlowsRow
	for rows.Next() {
		var i ReadAllFlowsRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Status,
			&i.ContainerID,
			&i.ContainerName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const readFlow = `-- name: ReadFlow :one
SELECT
  f.id, f.created_at, f.updated_at, f.name, f.status, f.container_id,
  c.name AS container_name,
  c.image AS container_image
FROM flows f
LEFT JOIN containers c ON f.container_id = c.id
WHERE f.id = $1
`

type ReadFlowRow struct {
	ID             int64
	CreatedAt      pgtype.Timestamp
	UpdatedAt      pgtype.Timestamp
	Name           pgtype.Text
	Status         pgtype.Text
	ContainerID    pgtype.Int8
	ContainerName  pgtype.Text
	ContainerImage pgtype.Text
}

func (q *Queries) ReadFlow(ctx context.Context, id int64) (ReadFlowRow, error) {
	row := q.db.QueryRow(ctx, readFlow, id)
	var i ReadFlowRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Status,
		&i.ContainerID,
		&i.ContainerName,
		&i.ContainerImage,
	)
	return i, err
}

const updateFlowName = `-- name: UpdateFlowName :one
UPDATE flows
SET name = $1
WHERE id = $2
RETURNING id, created_at, updated_at, name, status, container_id
`

type UpdateFlowNameParams struct {
	Name pgtype.Text
	ID   int64
}

func (q *Queries) UpdateFlowName(ctx context.Context, arg UpdateFlowNameParams) (Flow, error) {
	row := q.db.QueryRow(ctx, updateFlowName, arg.Name, arg.ID)
	var i Flow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Status,
		&i.ContainerID,
	)
	return i, err
}

const updateFlowStatus = `-- name: UpdateFlowStatus :one
UPDATE flows
SET status = $1
WHERE id = $2
RETURNING id, created_at, updated_at, name, status, container_id
`

type UpdateFlowStatusParams struct {
	Status pgtype.Text
	ID     int64
}

func (q *Queries) UpdateFlowStatus(ctx context.Context, arg UpdateFlowStatusParams) (Flow, error) {
	row := q.db.QueryRow(ctx, updateFlowStatus, arg.Status, arg.ID)
	var i Flow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Status,
		&i.ContainerID,
	)
	return i, err
}

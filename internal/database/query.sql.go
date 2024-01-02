// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: query.sql

package database

import (
	"context"
	"database/sql"
)

const getCategoryTree = `-- name: GetCategoryTree :many
WITH RECURSIVE cte_categories (id, name, parent_category_id) AS (
    SELECT c.id, c.name, c.parent_category_id
    FROM categories c
    JOIN products p ON c.id = p.category_id
    WHERE p.id = ?

    UNION ALL

    SELECT c.id, c.name, c.parent_category_id
    FROM categories c
    JOIN cte_categories ct ON c.id = ct.parent_category_id
)

SELECT id, name, parent_category_id FROM cte_categories
`

type GetCategoryTreeRow struct {
	ID               int64
	Name             string
	ParentCategoryID sql.NullInt64
}

func (q *Queries) GetCategoryTree(ctx context.Context, id string) ([]GetCategoryTreeRow, error) {
	rows, err := q.db.QueryContext(ctx, getCategoryTree, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCategoryTreeRow
	for rows.Next() {
		var i GetCategoryTreeRow
		if err := rows.Scan(&i.ID, &i.Name, &i.ParentCategoryID); err != nil {
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

const listProducts = `-- name: ListProducts :many
SELECT id, name, category_id FROM products
ORDER BY name
`

func (q *Queries) ListProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, listProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(&i.ID, &i.Name, &i.CategoryID); err != nil {
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

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package database

import (
	"database/sql"
)

type Category struct {
	ID               int64
	Name             string
	ParentCategoryID sql.NullInt64
}

type Product struct {
	ID         int64
	Name       string
	CategoryID sql.NullInt64
}

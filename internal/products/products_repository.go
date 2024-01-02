package products

import (
	"context"
	"database/sql"

	"github.com/soda-zero/vegandb/internal/database"
)

type ProductRepository struct {
	queries *database.Queries
}

func NewProductRepository(queries *database.Queries) *ProductRepository {
	return &ProductRepository{queries: queries}
}

// ListProducts retrieves a list of all products.
func (r *ProductRepository) ListProducts(ctx context.Context) ([]database.Product, error) {
	return r.queries.ListProducts(ctx)
}

// GetCategoryTreeForProduct retrieves the category tree for a given product ID.
func (r *ProductRepository) GetCategoryTreeForProduct(ctx context.Context, productID int64) ([]database.GetCategoryTreeRow, error) {
	return r.queries.GetCategoryTree(ctx, productID)
}

// GetFilterProducts retrieves the filtered products tree for a given product.
func (r *ProductRepository) FilterProducts(ctx context.Context, filter string) ([]database.Product, error) {
	nullFilter := sql.NullString{String: filter, Valid: filter != ""}
	return r.queries.FilterProducts(ctx, nullFilter)
}

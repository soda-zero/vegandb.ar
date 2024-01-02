package products

import (
	"context"

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
func (r *ProductRepository) GetCategoryTreeForProduct(ctx context.Context, productID string) ([]database.GetCategoryTreeRow, error) {
	return r.queries.GetCategoryTree(ctx, productID)
}

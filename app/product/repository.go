package product

import "context"

//go:generate mockgen -source=repository.go -destination=mock_repository.go -package=product
type Repository interface {
	GetProductByID(ctx context.Context, id string) (*Product, error)
	GetProductsByIDs(ctx context.Context, ids []string) ([]Product, error)
	CreateProduct(ctx context.Context, product *Product) (*Product, error)
}

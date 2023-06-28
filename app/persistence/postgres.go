package persistence

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/pact-cdc-example/product-service/app/product"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -source=postgres.go -destination=mock_postgres_repository.go -package=persistence
type PostgresRepository interface {
	GetProductByID(ctx context.Context, id string) (*product.Product, error)
	GetProductsByIDs(ctx context.Context, ids []string) ([]product.Product, error)
	CreateProduct(ctx context.Context, product *product.Product) (*product.Product, error)
}

type postgresRepository struct {
	db     *sql.DB
	logger *logrus.Logger
}

type NewPostgresRepositoryOpts struct {
	DB *sql.DB
	L  *logrus.Logger
}

func NewPostgresRepository(opts *NewPostgresRepositoryOpts) PostgresRepository {
	return &postgresRepository{
		db:     opts.DB,
		logger: opts.L,
	}
}

func (pr *postgresRepository) GetProductByID(
	ctx context.Context, id string) (*product.Product, error) {
	row := pr.db.QueryRowContext(
		ctx,
		`SELECT id, name, code, color, created_at, updated_at, 
    buying_price, selling_price, image_url, type, provider, creator,
    distributor
		FROM products WHERE ID = $1`,
		id,
	)

	var p product.Product
	if err := row.Scan(
		&p.ID,
		&p.Name,
		&p.Code,
		&p.Color,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.BuyingPrice,
		&p.SellingPrice,
		&p.ImageURL,
		&p.Type,
		&p.Provider,
		&p.Creator,
		&p.Distributor,
	); err != nil {
		pr.logger.Errorf("could not scan product :%v", err)
		return nil, err
	}

	return &p, nil
}

func (pr *postgresRepository) GetProductsByIDs(
	ctx context.Context, ids []string) ([]product.Product, error) {
	products := make([]product.Product, 0, len(ids))
	for i := range ids {
		product, err := pr.GetProductByID(ctx, ids[i])
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		if product != nil {
			products = append(products, *product)
		}
	}

	return products, nil
}

func (pr *postgresRepository) CreateProduct(
	ctx context.Context, product *product.Product) (*product.Product, error) {
	row := pr.db.QueryRowContext(
		ctx,
		`INSERT INTO products (id, name, code, color, buying_price, selling_price,
		image_url, type, provider, creator, distributor)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING created_at, updated_at`,
		product.ID,
		product.Name,
		product.Code,
		product.Color,
		product.BuyingPrice,
		product.SellingPrice,
		product.ImageURL,
		product.Type,
		product.Provider,
		product.Creator,
		product.Distributor,
	)

	var createdAt time.Time
	var updatedAt time.Time

	if err := row.Scan(&createdAt, &updatedAt); err != nil {
		pr.logger.Errorf("could not get created product :%v", err)
		return nil, err
	}

	product.CreatedAt = createdAt
	product.UpdatedAt = updatedAt

	return product, nil
}

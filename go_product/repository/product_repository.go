package repository

import (
	"context"
	"database/sql"

	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/model/domain"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, Product domain.Product) domain.Product
	Update(ctx context.Context, tx *sql.Tx, Product domain.Product) domain.Product
	Delete(ctx context.Context, tx *sql.Tx, Product domain.Product) domain.Product
	FindById(ctx context.Context, tx *sql.Tx, productKode string) (domain.Product, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Product
}

package repository

import (
	"context"
	"database/sql"

	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/model/domain"
	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/model/web"
)

type ProductCartRepository interface {
	Save(ctx context.Context, tx *sql.Tx, ProductCart domain.ProductCart) domain.ProductCart
	Update(ctx context.Context, tx *sql.Tx, ProductCart domain.ProductCart) domain.ProductCart
	Delete(ctx context.Context, tx *sql.Tx, ProductCart domain.ProductCart) domain.ProductCart
	FindById(ctx context.Context, tx *sql.Tx, kodeProduk string) (domain.ProductCart, error)
	FindAll(ctx context.Context, tx *sql.Tx) []web.ProductCartResponse
}

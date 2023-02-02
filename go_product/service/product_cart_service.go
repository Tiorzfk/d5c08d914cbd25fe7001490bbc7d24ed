package service

import (
	"context"

	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/model/web"
)

type ProductCartService interface {
	Create(ctx context.Context, request web.ProductCartCreateRequest) web.ProductCartResponse
	Delete(ctx context.Context, kode string) web.ProductCartResponse
	FindById(ctx context.Context, kode string) web.ProductCartResponse
	FindAll(ctx context.Context) []web.ProductCartResponse
}

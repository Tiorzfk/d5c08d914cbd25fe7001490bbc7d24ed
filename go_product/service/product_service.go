package service

import (
	"context"

	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/model/web"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse
	Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse
	Delete(ctx context.Context, kode string) web.ProductResponse
	FindById(ctx context.Context, kode string) web.ProductResponse
	FindAll(ctx context.Context) []web.ProductResponse
}

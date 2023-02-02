package service

import (
	"context"
	"database/sql"

	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/exception"
	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/helper"
	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/model/domain"
	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/model/web"
	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/repository"

	"github.com/go-playground/validator"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, DB *sql.DB, validate *validator.Validate) *ProductServiceImpl {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product := domain.Product{
		Kode: request.Kode,
		Nama: request.Nama,
		Stok: request.Stok,
	}

	product = service.ProductRepository.Save(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, request.Kode)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	product.Kode = request.Kode
	product.Nama = request.Nama
	product.Stok = request.Stok

	product = service.ProductRepository.Update(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) Delete(ctx context.Context, kode string) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, kode)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	product = service.ProductRepository.Delete(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) FindById(ctx context.Context, kode string) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err2 := service.ProductRepository.FindById(ctx, tx, kode)
	if err2 != nil {
		panic(exception.NewNotFoundError(err2.Error()))
	}

	return helper.ToProductResponse(category)
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) []web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := service.ProductRepository.FindAll(ctx, tx)
	helper.PanicIfError(err)

	var productResponses []web.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, helper.ToProductResponse(product))
	}

	return productResponses
}

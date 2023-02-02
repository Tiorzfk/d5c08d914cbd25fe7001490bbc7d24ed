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

type ProductCartServiceImpl struct {
	ProductCartRepository repository.ProductCartRepository
	DB                    *sql.DB
	Validate              *validator.Validate
}

func NewProductCartService(productCartRepository repository.ProductCartRepository, DB *sql.DB, validate *validator.Validate) *ProductCartServiceImpl {
	return &ProductCartServiceImpl{
		ProductCartRepository: productCartRepository,
		DB:                    DB,
		Validate:              validate,
	}
}

func (service *ProductCartServiceImpl) Create(ctx context.Context, request web.ProductCartCreateRequest) web.ProductCartResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	cart, err := service.ProductCartRepository.FindById(ctx, tx, request.KodeProduk)
	product := domain.ProductCart{
		KodeProduk: request.KodeProduk,
		Kuantitas:  request.Kuantitas,
	}

	if err != nil {
		product = service.ProductCartRepository.Save(ctx, tx, product)
	} else {
		product = domain.ProductCart{
			KodeProduk: request.KodeProduk,
			Kuantitas:  cart.Kuantitas + 1,
		}
		product = service.ProductCartRepository.Update(ctx, tx, product)
	}

	return helper.ToProductCartResponse(product)
}

func (service *ProductCartServiceImpl) Delete(ctx context.Context, kode string) web.ProductCartResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductCartRepository.FindById(ctx, tx, kode)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	product = service.ProductCartRepository.Delete(ctx, tx, product)

	return helper.ToProductCartResponse(product)
}

func (service *ProductCartServiceImpl) FindById(ctx context.Context, kode string) web.ProductCartResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err2 := service.ProductCartRepository.FindById(ctx, tx, kode)
	if err2 != nil {
		panic(exception.NewNotFoundError(err2.Error()))
	}

	return helper.ToProductCartResponse(product)
}

func (service *ProductCartServiceImpl) FindAll(ctx context.Context) []web.ProductCartResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := service.ProductCartRepository.FindAll(ctx, tx)
	helper.PanicIfError(err)

	var productCartResponses []web.ProductCartResponse
	for _, product := range products {
		productCartResponses = append(productCartResponses, product)
	}

	return productCartResponses
}

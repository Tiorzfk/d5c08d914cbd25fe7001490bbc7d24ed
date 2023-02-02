package main

import (
	"net/http"

	"github.com/go-playground/validator"

	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/app"
	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/controller"
	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/helper"
	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/repository"
	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/service"
)

func main() {
	productRepositoryImpl := repository.NewProductRepository()
	productCartRepositoryImpl := repository.NewProductCartRepository()
	db := app.NewDB()
	validate := validator.New()
	productServiceImpl := service.NewProductService(productRepositoryImpl, db, validate)
	productCartServiceImpl := service.NewProductCartService(productCartRepositoryImpl, db, validate)
	productControllerImpl := controller.NewProductController(productServiceImpl)
	productCartControllerImpl := controller.NewProductCartController(productCartServiceImpl)

	router := app.NewRouter(productControllerImpl, productCartControllerImpl)
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

package app

import (
	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/controller"
	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/exception"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(produkController controller.ProductController, produkCartController controller.ProductCartController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/products", produkController.FindAll)
	router.GET("/api/products/:kode", produkController.FindById)
	router.POST("/api/products", produkController.Create)
	router.PUT("/api/products/:kode", produkController.Update)
	router.DELETE("/api/products/:kode", produkController.Delete)

	router.GET("/api/cart/:kode", produkCartController.FindAll)
	router.POST("/api/cart", produkCartController.Create)

	router.PanicHandler = exception.ErrorHandler

	return router
}

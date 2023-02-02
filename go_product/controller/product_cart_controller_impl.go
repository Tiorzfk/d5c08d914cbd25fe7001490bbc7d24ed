package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/helper"
	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/model/web"
	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/service"

	"github.com/julienschmidt/httprouter"
)

type ProductCartControllerImpl struct {
	ProductCartService service.ProductCartService
}

func NewProductCartController(productCartService service.ProductCartService) *ProductCartControllerImpl {
	return &ProductCartControllerImpl{
		ProductCartService: productCartService,
	}
}

func (controller *ProductCartControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productCartCreateRequest := web.ProductCartCreateRequest{}
	helper.ReadFromRequestBody(request, &productCartCreateRequest)
	productCartResponse := controller.ProductCartService.Create(request.Context(), productCartCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productCartResponse,
	}

	writer.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(writer)
	errEncode := encoder.Encode(webResponse)
	helper.PanicIfError(errEncode)
}

func (controller *ProductCartControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kode := params.ByName("kode")

	productCartResponse := controller.ProductCartService.Delete(request.Context(), kode)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productCartResponse,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *ProductCartControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productCartResponses := controller.ProductCartService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productCartResponses,
	}

	helper.WriteResponseBody(writer, webResponse)
}

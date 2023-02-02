package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/helper"
	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/model/web"
	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/service"

	"github.com/julienschmidt/httprouter"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductControllerImpl {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller *ProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productCreateRequest := web.ProductCreateRequest{}
	helper.ReadFromRequestBody(request, &productCreateRequest)
	productResponse := controller.ProductService.Create(request.Context(), productCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	writer.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(writer)
	errEncode := encoder.Encode(webResponse)
	helper.PanicIfError(errEncode)
}
func (controller *ProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productUpdateRequest := web.ProductUpdateRequest{}
	helper.ReadFromRequestBody(request, &productUpdateRequest)

	kode := params.ByName("kode")
	_, err := strconv.Atoi(kode)
	helper.PanicIfError(err)

	productResponse := controller.ProductService.Update(request.Context(), productUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteResponseBody(writer, webResponse)
}
func (controller *ProductControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kode := params.ByName("kode")

	productResponse := controller.ProductService.Delete(request.Context(), kode)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kode := params.ByName("kode")
	productResponse := controller.ProductService.FindById(request.Context(), kode)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteResponseBody(writer, webResponse)
}
func (controller *ProductControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productResponses := controller.ProductService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponses,
	}

	helper.WriteResponseBody(writer, webResponse)
}

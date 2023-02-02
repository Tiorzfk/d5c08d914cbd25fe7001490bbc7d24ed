package exception

import (
	"net/http"

	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/helper"
	"github.com/Tiorzfk/d5c08d914cbd25fe7001490bbc7d24ed/model/web"

	"github.com/go-playground/validator"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writer, request, err) {
		return
	}

	if validationErrors(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webReponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error(),
		}

		helper.WriteResponseBody(writer, webReponse)
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webReponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}

		helper.WriteResponseBody(writer, webReponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webReponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	helper.WriteResponseBody(writer, webReponse)
}

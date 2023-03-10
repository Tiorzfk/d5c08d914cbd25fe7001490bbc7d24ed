package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&result)
	PanicIfError(err)
}

func WriteResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(writer)
	errEncode := encoder.Encode(response)
	PanicIfError(errEncode)
}

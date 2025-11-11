package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	contentTypeHeader = "Content-Type"
	jsonContentType   = "application/json; charset=UTF-8"
)

func JSONResponse(writer http.ResponseWriter, response any, code int) {
	writer.WriteHeader(code)
	writer.Header().Set(contentTypeHeader, jsonContentType)

	if err := json.NewEncoder(writer).Encode(response); err != nil {
		log.Println(fmt.Errorf("json.Encode: %w", err).Error())
	}
}

func JSONBaseError(writer http.ResponseWriter, message string, code int) {
	response := errorResponseDTO{
		Error: errorDTO{
			Message: message,
		},
	}
	JSONResponse(writer, response, code)
}

func JSONBadRequestError(writer http.ResponseWriter) {
	JSONBaseError(writer, "Bad Request", http.StatusBadRequest)
}

func JSONInternalServerError(writer http.ResponseWriter) {
	JSONBaseError(writer, "Internal Server Error", http.StatusInternalServerError)
}

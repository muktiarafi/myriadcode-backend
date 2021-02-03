package helpers

import (
	"encoding/json"
	"net/http"
)

type BaseResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	BaseResponse
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	BaseResponse
	Error string `json:"error"`
}

func SendJSON(w http.ResponseWriter, status int, message string, data interface{}) error {
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(&SuccessResponse{
		BaseResponse{status, message},
		data,
	})
}

func SendError(w http.ResponseWriter, status int, message, error string) error {
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(&ErrorResponse{
		BaseResponse{status, message},
		error,
	})
}


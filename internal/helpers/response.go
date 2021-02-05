package helpers

import (
	"encoding/json"
	"github.com/muktiarafi/myriadcode-backend/internal/apierror"
	"github.com/muktiarafi/myriadcode-backend/internal/configs"
	"net/http"
	"runtime/debug"
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

var app *configs.AppConfig

func NewHelper(a *configs.AppConfig) {
	app = a
}

func SendJSON(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(&SuccessResponse{
		BaseResponse{status, "OK"},
		data,
	})
}

func SendError(w http.ResponseWriter, err error) {

	if e, ok := err.(*apierror.Error); ok {
		w.WriteHeader(e.StatusCode)

		json.NewEncoder(w).Encode(&ErrorResponse{
			BaseResponse{e.StatusCode, e.Kind.Error()},
			e.Warning,
		})
		app.WarningLog.Printf("Client Error: %s", e.Error())
	} else {
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(&ErrorResponse{
			BaseResponse{http.StatusInternalServerError, apierror.InternalServerError.Error()},
			apierror.InternalServerError.Error(),
		})
		app.ErrorLog.Printf("Server Error!!!: %s, stack trace:\n%s", err.Error(), debug.Stack())
	}
}

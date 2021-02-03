package handler

import (
	"github.com/muktiarafi/myriadcode-backend/internal/apierror"
	"github.com/muktiarafi/myriadcode-backend/internal/configs"
	"github.com/muktiarafi/myriadcode-backend/internal/helpers"
	"net/http"
	"runtime/debug"
)

var app *configs.AppConfig

// InitializeHandlerConfig initialize app so the custom handler can use embedded custom logger for logging
func InitializeHandlerConfig(a *configs.AppConfig) {
	app = a
}

// CustomHandler make it possible to return error from handlerfunc
type CustomHandler func(w http.ResponseWriter, r *http.Request) error

// ServeHTTP is method implementation for http handler
func (ch CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := ch(w, r)
	if err != nil {
		switch e := err.(type) {
		case *apierror.Error:
			helpers.SendError(
				w,
				e.StatusCode,
				http.StatusText(e.StatusCode),
				e.Kind.Error(),
			)
			app.WarningLog.Println("Masih santuy Error")
		default:
			helpers.SendError(
				w,
				http.StatusInternalServerError,
				apierror.InternalServerError.Error(),
				"Server Error, Try Again later",
			)
			// need to rework this because stack trace can't log wrapped handlerfunc
			// and return trace from this handler instead
			app.ErrorLog.Printf("Error with stack trace: \n %s", debug.Stack())
		}
	}
}

package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/muktiarafi/myriadcode-backend/internal/controllers"
	"github.com/muktiarafi/myriadcode-backend/internal/handler"
	"net/http"
)

func SetRouter() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.SetHeader("Content-Type", "application/json"))
	mux.Method(http.MethodGet, "/", handler.CustomHandler(controllers.Test))

	return mux
}

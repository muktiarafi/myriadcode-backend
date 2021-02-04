package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func SetRouter() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.SetHeader("Content-Type", "application/json"))

	return mux
}

package routes

import (
	"github.com/dami-pie/api/controller"
	"github.com/go-chi/chi/v5"
)

func AddRoutes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", controller.TestHandler)
	router.Get("/login", controller.LoginHandler)

	return router
}

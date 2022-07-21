package router

import (
	"github.com/dami-pie/api/server/router/route"
	"github.com/gorilla/mux"
)

func AddRoutes() *mux.Router {
	r := mux.NewRouter()
	return route.ConfigRoutes(r)
}

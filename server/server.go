package server

import (
	"net/http"

	"github.com/dami-pie/api/server/routes"
)

func Run() {
	http.ListenAndServe(":3000", routes.AddRoutes())
}

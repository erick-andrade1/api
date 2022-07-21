package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar coloca todas as rotas dentro do router
func ConfigRoutes(r *mux.Router) *mux.Router {
	routes := routeGoogleAuth

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Funcao).Methods(route.Metodo)
	}

	return r
}

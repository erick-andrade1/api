package route

import (
	"net/http"

	"github.com/dami-pie/api/controller"
)

var routeGoogleAuth = []Route{
	{
		URI:                "/",
		Metodo:             http.MethodGet,
		Funcao:             controller.IndexRoute,
		RequerAutenticacao: false,
	},
	{
		URI:                "/auth/{provider}/callback",
		Metodo:             http.MethodGet,
		Funcao:             controller.LoginHandler,
		RequerAutenticacao: false,
	},
	{
		URI:                "/auth/{provider}",
		Metodo:             http.MethodGet,
		Funcao:             controller.AuthHandler,
		RequerAutenticacao: false,
	},
}

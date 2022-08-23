package route

import (
	"net/http"

	"github.com/dami-pie/api/controller"
)

var apiRoutes = []Route{
	{
		URI:                "/authenticate",
		Metodo:             http.MethodPost,
		Funcao:             controller.Login,
		RequerAutenticacao: false,
	},
	{
		URI:                "/",
		Metodo:             http.MethodPost,
		Funcao:             controller.AuthOTP,
		RequerAutenticacao: true,
	},
}

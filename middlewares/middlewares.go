package middlewares

import (
	"fmt"
	"net/http"

	"github.com/dami-pie/api/middlewares/auth"
	"github.com/dami-pie/api/responses"
)

// Escreve as informações da requisição no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Printf("\n %s %s %s", req.Method, req.RequestURI, req.Host)
		res.Header().Add("Content-Type", "application/json")
		next(res, req)
	}
}

// Verifica se o usuário está autenticado, se estiver ele irá para próxima função
func AuthenticateUser(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if erro := auth.ValidateToken(req); erro != nil {
			responses.Erro(res, http.StatusUnauthorized, erro)
		} else {
			next(res, req)
		}

	}
}

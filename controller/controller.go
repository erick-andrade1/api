package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dami-pie/api/middlewares/auth"
	"github.com/dami-pie/api/models"
	"github.com/dami-pie/api/responses"
)

var token struct {
	Token string
}

func Login(res http.ResponseWriter, req *http.Request) {
	// Lê o body da requisição
	body, erro := ioutil.ReadAll(req.Body)
	if body == nil {
		responses.Erro(res, http.StatusUnprocessableEntity, erro)
		return
	}

	// Descompacta o JSON recebido
	var usuario models.User
	if erro = json.Unmarshal(body, &usuario); erro != nil {
		responses.Erro(res, http.StatusBadRequest, erro)
		return
	}

	// Valida o usuário
	if erro = usuario.PrepareUser(); erro != nil {
		responses.Erro(res, http.StatusBadRequest, erro)
		return
	}

	// Gera o token
	token, erro := auth.GenerateToken(usuario.Email)
	if erro != nil {
		responses.Erro(res, http.StatusInternalServerError, erro)
		return
	}

	// Devolve o token gerado
	responses.JSON(res, http.StatusAccepted, token)
}

func AuthOTP(res http.ResponseWriter, req *http.Request) {
	email, erro := auth.ExtractUserEmail(req)
	if erro != nil {
		responses.Erro(res, http.StatusUnauthorized, erro)
	}

	fmt.Println("\n Email: ", email)
}

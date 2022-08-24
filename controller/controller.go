package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dami-pie/api/config"
	"github.com/dami-pie/api/middlewares/auth"
	"github.com/dami-pie/api/models"
	"github.com/dami-pie/api/responses"
)

type jsonToken struct {
	Token string `json:"token"`
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

	// Apenas formata o json bonitinho:
	userToken := jsonToken{Token: token}

	// Devolve o token gerado
	responses.JSON(res, http.StatusAccepted, userToken)
}

func AuthOTP(res http.ResponseWriter, req *http.Request) {
	_, erro := auth.ExtractUserEmail(req)
	if erro != nil {
		responses.Erro(res, http.StatusUnauthorized, erro)
	}

	body, erro := ioutil.ReadAll(req.Body)
	if body == nil {
		responses.Erro(res, http.StatusUnprocessableEntity, erro)
		return
	}

	var otp models.OTP
	if erro := json.Unmarshal(body, &otp); erro != nil {
		responses.Erro(res, http.StatusUnprocessableEntity, erro)
		return
	}

	if isValid, erro := otp.ValidateKey(); !isValid {
		responses.Erro(res, http.StatusUnauthorized, erro)
		return
	} else {
		sendOpenCommand()
	}

	//TODO: Precisamos guardar o email e o tempo da entrada do usuário. Isso vai ser implementado junto com o banco.
}

func sendOpenCommand() {
	body, _ := json.Marshal(map[string]any{
		"hash": config.OTPKey,
		"open": true,
	})
	payload := bytes.NewBuffer(body)
	http.Post(config.Device, "application/json", payload)
}

package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dami-pie/api/config"
	"github.com/dgrijalva/jwt-go/v4"
)

// Função responsável por gerar o token jwt
func GenerateToken(userEmail string) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["auth"] = true
	permissions["expiresAt"] = time.Now().Add(time.Hour * 1).Unix()
	permissions["email"] = userEmail

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(config.SecretKey))
}

func ValidateToken(req *http.Request) error {
	// é preciso dar um parse na string para que seja possível pegar os valores da permissão
	tokenString := extractToken(req)
	token, erro := jwt.Parse(tokenString, returnVerificationKey)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Token inválido!")
}

// Extrai o Email do usuário que está salvo com o Token
func ExtractUserEmail(req *http.Request) (string, error) {
	tokenString := extractToken(req)
	token, erro := jwt.Parse(tokenString, returnVerificationKey)
	if erro != nil {
		return "", erro
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Como o parametro permissoes["email"] não é uma string, é preciso transformar ele em string
		userEmail := permissoes["email"]
		if str, ok := userEmail.(string); ok {
			return str, nil
		} else {
			return "", errors.New("Email não reconhecido!")
		}
	}

	return "", errors.New("Token inválido!")
}

// Extrai o token da requisição "Bearer 'token'"
func extractToken(req *http.Request) string {
	token := req.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}

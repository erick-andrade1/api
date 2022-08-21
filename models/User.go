package models

import (
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

// O 'omitempty' irá fazer com que caso o campo ID esteja em branco quando for passar para JSON, o campo id não será passado no json
type User struct {
	Email string `json:"email"`
}

// Irá validar e formatar o usuário antes de inserir no banco de dados:
func (user *User) PrepareUser() error {
	if erro := user.validate(); erro != nil {
		return erro
	}

	user.format()
	return nil
}

func (user *User) validate() error {
	i := strings.LastIndexByte(user.Email, '@')
	host := user.Email[i+1:]

	// Se i < 0 não teve input de email.
	if i < 0 {
		return errors.New("O email é obrigatório e não pode estar em branco!")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("O email inserido é inválido")
	}

	if host != "ecomp.poli.br" {
		return errors.New("Host inválido!!")
	}

	return nil
}

func (user *User) format() {
	user.Email = strings.TrimSpace(user.Email)
}

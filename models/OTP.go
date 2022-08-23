package models

import (
	"errors"
	"time"

	"github.com/dami-pie/api/config"
	"github.com/pquerna/otp/totp"
)

type OTP struct {
	Tempo time.Time `json:"time"`
	Key   string    `json:"key"`
}

func (otp *OTP) ValidateKey() (bool, error) {
	if totp.Validate(otp.Key, config.OTPKey) {
		return true, nil
	} else {
		return false, errors.New("OTP inválida")
	}
}

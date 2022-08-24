package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	StringConexaoBanco = ""
	Port               = ""
	JWTKey             []byte
	OTPKey             = ""
	Device             = ""
)

func LoadApiConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Port = os.Getenv("PORT")
	JWTKey = []byte(os.Getenv("JWT_SECRET"))
	OTPKey = os.Getenv("OTP_SECRET")
	Device = os.Getenv("DEVICE")
}

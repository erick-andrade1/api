package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	StringConexaoBanco = ""
	Port               = ""
	SecretKey          []byte
)

func LoadApiConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Port = os.Getenv("PORT")
	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}

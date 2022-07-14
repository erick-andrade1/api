package server

import (
	"log"
	"net/http"
	"os"

	"github.com/dami-pie/api/server/routes"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

type Server struct {
	port   string
	routes chi.Router
}

func Run() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	server := Server{port, routes.AddRoutes()}
	http.ListenAndServe(server.port, server.routes)
}

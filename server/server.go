package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dami-pie/api/config"
	"github.com/dami-pie/api/server/router"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Server struct {
	port   string
	routes *mux.Router
}

func Run() {
	config.LoadApiConfig()

	server := Server{config.Port, router.AddRoutes()}

	fmt.Println("Server running on PORT:", server.port)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins
	})

	log.Fatal(http.ListenAndServeTLS(server.port, "certs/fullchain.pem", "certs/privkey.pem", c.Handler(server.routes)))
}

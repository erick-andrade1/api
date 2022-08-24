package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dami-pie/api/config"
	"github.com/dami-pie/api/server/router"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	port   string
	routes *mux.Router
}

func Run() {
	config.LoadApiConfig()

	server := Server{config.Port, router.AddRoutes()}

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("Server running on PORT:", server.port)

	log.Fatal(http.ListenAndServeTLS(server.port, "certs/fullchain.pem", "certs/privkey.pem", handlers.CORS(headers, methods, origins)(server.routes)))
}

package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dami-pie/api/config"
	"github.com/dami-pie/api/server/router"
	"github.com/gorilla/mux"
)

type Server struct {
	port   string
	routes *mux.Router
}

func Run() {
	config.LoadApiConfig()

	server := Server{config.Port, router.AddRoutes()}

	fmt.Println("Server running on PORT:", server.port)
	log.Fatal(http.ListenAndServe(server.port, server.routes))
}

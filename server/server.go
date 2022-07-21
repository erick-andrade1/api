package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dami-pie/api/server/router"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

type Server struct {
	port   string
	routes *mux.Router
}

func Run() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	var googleClientID string = os.Getenv("GOOGLE_CLIENT_ID")
	var googleClientSecret string = os.Getenv("GOOGLE_CLIENT_SECRET")

	server := Server{port, router.AddRoutes()}

	key := "Secret-session-key" // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30        // 30 days
	isProd := false             // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(
		google.New(googleClientID, googleClientSecret, "http://localhost:3000/auth/google/callback", "email", "profile"),
	)

	fmt.Println("Server running on PORT:", server.port)
	log.Fatal(http.ListenAndServe(server.port, server.routes))
}

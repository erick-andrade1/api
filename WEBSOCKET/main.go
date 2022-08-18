package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type userExample struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func reader(conn *websocket.Conn) {
	// Exemplo do websocket lendo algo recebido do front:
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func sender(conn *websocket.Conn) {
	// Exemplo do websocket enviando algo para o front:
	ch := time.Tick(5 * time.Second)
	for range ch {
		conn.WriteJSON(userExample{
			Username: "Basilisk Aeron",
			Name:     "Erick",
			LastName: "Veríssimo",
		})
	}
}

func websocketGet(w http.ResponseWriter, r *http.Request) {

	// Permite qualquer conexão com o endpoint websocket
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// "upando" a conexão para websocket
	websocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client connected....")

	go sender(websocket)
}

func websocketPost(w http.ResponseWriter, r *http.Request) {
	// Permite qualquer conexão com o endpoint websocket
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// "upando" a conexão para websocket
	websocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client connected....")

	reader(websocket)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws/post", websocketPost)
	http.HandleFunc("/ws/get", websocketGet)
}

func main() {
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8001", nil))
}

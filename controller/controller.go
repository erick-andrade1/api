package controller

import "net/http"

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, logged in!"))
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, testing!"))
}

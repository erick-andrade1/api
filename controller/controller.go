package controller

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/markbates/goth/gothic"
)

func LoginHandler(res http.ResponseWriter, req *http.Request) {
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}
	t, _ := template.ParseFiles("templates/success.html")
	t.Execute(res, user)
}

func AuthHandler(res http.ResponseWriter, req *http.Request) {
	gothic.BeginAuthHandler(res, req)
}

func IndexRoute(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(res, false)
}

package controllers

import (
	"net/http"
)

//top画面に飛ばす
func signup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "Hello", "layout", "signup")
}

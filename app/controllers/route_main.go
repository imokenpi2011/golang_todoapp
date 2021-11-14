package controllers

import (
	"net/http"
)

//top画面に飛ばす
func top(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "Hello", "layout", "top")
}

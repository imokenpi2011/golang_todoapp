package controllers

import (
	"net/http"
	"text/template"
)

//top画面に飛ばす
func top(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("app/views/templates/top.html")
	t.Execute(w, nil)
}

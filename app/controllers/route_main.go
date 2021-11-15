package controllers

import (
	"net/http"
)

//top画面に飛ばす
func top(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "Hello", "layout", "public_navbar", "top")
}

//タスク表示画面の制御
func index(w http.ResponseWriter, r *http.Request) {
	//セッションを取得する
	_, err := session(w, r)
	if err != nil {
		//セッションが存在しない場合はトップページに遷移する
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		//セッションが存在する場合はタスク表示画面に遷移する
		generateHTML(w, nil, "layout", "private_navbar", "index")
	}
}

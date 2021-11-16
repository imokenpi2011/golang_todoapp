package controllers

import (
	"log"
	"net/http"
)

//top画面に飛ばす
func top(w http.ResponseWriter, r *http.Request) {
	//セッションを取得する
	_, err := session(w, r)
	if err != nil {
		//セッションが存在しない場合はトップページに遷移する
		generateHTML(w, "Hello", "layout", "public_navbar", "top")
	} else {
		//セッションが存在する場合はタスク表示画面に遷移する
		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}

//タスク表示画面の制御
func index(w http.ResponseWriter, r *http.Request) {
	//セッションを取得する
	sess, err := session(w, r)
	if err != nil {
		//セッションが存在しない場合はトップページに遷移する
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		//ログイン中のユーザー情報を取得する
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		//ユーザーに紐づくタスクを取得する
		todos, _ := user.GetTodosByUser()
		user.Todos = todos
		//セッションが存在する場合はタスク表示画面に遷移する
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}

package controllers

import (
	"golang_todoapp/app/models"
	"log"
	"net/http"
)

//サインアップ画面の処理
func signup(w http.ResponseWriter, r *http.Request) {
	//GETの時に実行する
	if r.Method == "GET" {
		generateHTML(w, "Hello", "layout", "public_navbar", "signup")
	} else if r.Method == "POST" {
		//POSTの時に実行する
		//入力した値を受け取る
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		//インスタンスを生成
		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		//ユーザーを作成する
		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}

		//登録したらトップ画面に遷移する
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

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

//ログイン画面の処理
func login(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "Hello", "layout", "public_navbar", "login")
}

//ログインの検証
func authenticate(w http.ResponseWriter, r *http.Request) {
	//フォームの値を検証する
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	//フォームから値をもとにユーザーを検索する
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Panicln(err)
		http.Redirect(w, r, "/login", http.StatusFound)
	}

	//パスワードを検証する
	if user.PassWord == models.Encrypt(r.PostFormValue("password")) {
		//セッションを作成する
		err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}

		//作成したセッションを取得する
		session, err := models.GetSession(user.ID, user.Email)
		if err != nil {
			log.Println(err)
		}
		if err != nil {
			log.Println(err)
		}

		//クッキーを作成する
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)

		//ログインに成功したらトップに遷移する
		http.Redirect(w, r, "/", http.StatusFound)

	} else {
		//ログインに失敗した場合はログインページに遷移する
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

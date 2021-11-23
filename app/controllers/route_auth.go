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
		//セッションを取得する
		_, err := session(w, r)
		if err != nil {
			//セッションが存在しない場合はサインアップ画面に遷移する
			generateHTML(w, "Hello", "layout", "public_navbar", "signup")
		} else {
			//セッションが存在する場合はタスク表示画面に遷移する
			http.Redirect(w, r, "/todos", http.StatusFound)
		}
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
		if err := user.CreateUser(models.Db); err != nil {
			log.Println(err)
		}

		//登録したらトップ画面に遷移する
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

//ログイン画面の処理
func login(w http.ResponseWriter, r *http.Request) {
	//セッションを取得する
	_, err := session(w, r)
	if err != nil {
		//セッションが存在しない場合はログイン画面に遷移する
		generateHTML(w, "Hello", "layout", "public_navbar", "login")
	} else {
		//セッションが存在する場合はタスク表示画面に遷移する
		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}

//ログインの検証
func authenticate(w http.ResponseWriter, r *http.Request) {
	//フォームの値を検証する
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	//フォームから値をもとにユーザーを検索する
	user, err := models.GetUserByEmail(models.Db, r.PostFormValue("email"))
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

//ログアウトする
func logout(w http.ResponseWriter, r *http.Request) {
	//クッキーを取得
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}

	//クッキーの存在エラー以外の場合はセッションを削除する
	if err != http.ErrNoCookie {
		session := models.Session{UUID: cookie.Value}
		session.DeleteSessionByUUID()
	}
	http.Redirect(w, r, "/login", http.StatusFound)
}

package controllers

import (
	"golang_todoapp/app/models"
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

//タスク作成ページの遷移制御
func todoNew(w http.ResponseWriter, r *http.Request) {
	//セッションを取得する
	_, err := session(w, r)
	if err != nil {
		//セッションが存在しない場合はログインページに遷移する
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		//セッションが存在する場合はタスク作成画面に遷移する
		generateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}
}

//タスク作成処理
func todoSave(w http.ResponseWriter, r *http.Request) {
	//セッションを取得する
	sess, err := session(w, r)
	if err != nil {
		//セッションが存在しない場合はログインページに遷移する
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		//フォームの値を取得する
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		//ログイン中ユーザーの情報を取得する
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}

		//フォームから入力内容を受け取る
		content := r.PostFormValue("content")
		if err := user.CreateTodo(content); err != nil {
			log.Println(err)
		}

		//タスク一覧のページに遷移する
		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}

//タスク作成処理
func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	//セッションを取得する
	sess, err := session(w, r)
	if err != nil {
		//セッションが存在しない場合はログインページに遷移する
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		//セッションの確認
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		//編集対象のタスクを取得する
		t, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}

		//編集画面に遷移する
		generateHTML(w, t, "layout", "private_navbar", "todo_edit")
	}
}

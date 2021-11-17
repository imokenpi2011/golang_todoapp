package controllers

import (
	"fmt"
	"golang_todoapp/app/models"
	"golang_todoapp/config"
	"net/http"
	"regexp"
	"strconv"
	"text/template"
)

//ページを表示する
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	//html一覧を読み込む
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	//レイアウトを明示的に読み込む
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

//セッションを検証する
func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	//クッキーを取得する
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		//セッションを検証する
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("invalid session")
		}
	}
	return sess, err
}

//タスク編集・更新のURLがマッチしているか検証
var validPath = regexp.MustCompile("^/todos/(edit|update|delete)/([0-9]+$)")

//URLの検証
func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//タスク編集のURLとマッチしているか検証する
		// 例) /todis/edit/1
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			//マッチしない場合は404でリターンする
			http.NotFound(w, r)
			return
		}

		//URLのIDの部分が数字であるか検証
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			//マッチしない場合は404でリターンする
			http.NotFound(w, r)
			return
		}

		fn(w, r, qi)
	}
}

//サーバーを起動する
func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))

	//本来はstatic配下を読み込むが、StripPrefixでstaticを取り除いている
	http.Handle("/static/", http.StripPrefix("/static/", files))
	//トップページ。controller/route_main.goで管理
	//末尾に[/]がついていない場合は完全一致、そうでない場合は含まれていれば良い
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))

	//サーバを起動する。(起動ポート,関係ないページの場合404を返す様にする)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}

package controllers

import (
	"golang_todoapp/config"
	"net/http"
)

//サーバーを起動する
func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))

	//本来はstatic配下を読み込むが、StripPrefixでstaticを取り除いている
	http.Handle("/static/", http.StripPrefix("/static/", files))
	//トップページ。controller/route_main.goで管理
	http.HandleFunc("/", top)

	//サーバを起動する。(起動ポート,関係ないページの場合404を返す様にする)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}

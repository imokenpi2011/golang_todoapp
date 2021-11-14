package controllers

import (
	"golang_todoapp/config"
	"net/http"
)

//サーバーを起動する
func StartMainServer() error {
	//トップページ。controller/route_main.goで管理
	http.HandleFunc("/", top)

	//サーバを起動する。(起動ポート,関係ないページの場合404を返す様にする)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}

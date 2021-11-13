package models

import (
	"database/sql"
	"fmt"
	"golang_todoapp/config"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//テーブルの作成
var Db *sql.DB

var err error

const (
	tableNameUser = "users"
)

//初期処理
func init() {
	//sql.open(ドライバ名,DB名)でDBを読み込む
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	//存在しない場合はusersテーブルを作成する
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)

	//SQL実行
	Db.Exec(cmdU)
}

package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"golang_todoapp/config"
	"os"

	"github.com/google/uuid"
	"github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

//テーブルの作成
var Db *sql.DB

var err error

//初期処理
func init() {

	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	connection += "sslmode=require"
	Db, err = sql.Open(config.Config.SQLDriver, connection)
}

//UUIDを作成する
func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

//sha1を使用してハッシュ値を生成する
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}

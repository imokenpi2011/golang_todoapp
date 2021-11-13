package models

import (
	"log"
	"time"
)

type User struct {
	ID        int       //ユーザーID
	UUID      string    //UUID
	Name      string    //ユーザー名
	Email     string    //Eメール
	PassWord  string    //パスワード
	CreatedAt time.Time //作成日
}

//ユーザーを作成する
func (u *User) CreateUser() (err error) {
	//SQL文を指定
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

	//作成処理を実行
	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

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

//ユーザー情報を取得する
func GetUser(id int) (user User, err error) {
	//ユーザーインスタンスを宣言
	user = User{}

	//SQL文を設定
	cmd := `select id, uuid, name, email, password, created_at from users where id = ?`

	//取得処理を実行
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user, err
}

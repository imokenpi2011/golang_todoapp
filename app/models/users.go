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
	Todos     []Todo    //タスク一覧
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

//ユーザーを更新する
func (u *User) UpdateUser() (err error) {
	//SQL文を設定
	cmd := `update users set name = ?, email = ? where id = ?`

	//更新処理を実行
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

//ユーザーを削除する
func (u *User) DeleteUser() (err error) {
	//SQL文を作成
	cmd := `delete from users where id = ?`

	//削除処理を実行
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

//ログイン用ユーザー検索
func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at
	from users where email = ?`
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)

	return user, err
}

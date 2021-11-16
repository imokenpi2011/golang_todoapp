package models

import (
	"log"
	"time"
)

type Session struct {
	ID        int       //ユーザーID
	UUID      string    //UUID
	Email     string    //Eメール
	PassWord  string    //パスワード
	UserID    int       //ユーザーID
	CreatedAt time.Time //作成日
}

//セッションを作成する
func (u *User) CreateSession() (err error) {
	//SQL文を設定
	cmd1 := `insert into sessions (
		uuid,
		email,
		user_id,
		created_at) values (?,?,?,?)`

	//セッションの作成処理
	_, err = Db.Exec(cmd1, createUUID(), u.Email, u.ID, time.Now())
	if err != nil {
		log.Println(err)
	}

	return err
}

//生成したセッションを取得する
func GetSession(id int, email string) (session Session, err error) {
	//セッションインスタンスを宣言
	session = Session{}

	//SQL文を設定
	cmd := `select id, uuid, email, user_id, created_at
	from sessions where user_id = ? and email = ?`

	//セッション情報の取得
	err = Db.QueryRow(cmd, id, email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt)

	return session, err
}

//セッションをもとにユーザーを取得する
func (sess *Session) GetUserBySession() (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, created_at FROM users
	where id = ?`
	err = Db.QueryRow(cmd, sess.UserID).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.CreatedAt)

	return user, err
}

//セッションをUUIDで検証する
func (sess *Session) CheckSession() (valid bool, err error) {
	cmd := `select id, uuid, email, user_id, created_at
	from sessions where uuid = ?`

	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Email,
		&sess.UserID,
		&sess.CreatedAt)

	//セッションがある場合は処理を終了する
	if err != nil {
		valid = false
		return
	}

	//セッションがない場合はエラーを返却する
	if sess.ID != 0 {
		valid = true
	}
	return valid, err
}

//UUIDに一致するセッションを削除する。
func (sess *Session) DeleteSessionByUUID() (err error) {
	cmd := `delete from sessions where uuid = ?`
	_, err = Db.Exec(cmd, sess.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

package models

import (
	"log"
	"time"
)

type Todo struct {
	ID        int       //タスクID
	Content   string    //タスクの内容
	UserID    int       //ユーザーID
	CreatedAt time.Time //作成日時
}

//タスクを作成する
func (u *User) CreateTodo(content string) (err error) {
	//SQL文を設定
	cmd := `insert into todos (
		content,
		user_id,
		created_at) values (?,?,?)`

	//作成処理を実行
	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

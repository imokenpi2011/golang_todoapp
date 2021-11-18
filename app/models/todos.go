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
		created_at) values ($1,$2,$3)`

	//作成処理を実行
	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

//指定したIDのタスクを取得する
func GetTodo(id int) (todo Todo, err error) {
	todo = Todo{}

	//SQL文を指定
	cmd := `select id, content, user_id, created_at from todos where id = $1`

	//取得処理を実行
	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt,
	)
	return todo, err
}

//タスクの全一覧を取得する
func GetTodos() (todos []Todo, err error) {
	//SQL文を指定
	cmd := `select id, content, user_id, created_at from todos`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	//タスク情報の取得処理
	for rows.Next() {
		var todo Todo
		err := rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt,
		)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

//ユーザーに応じたタスク一覧を取得する
func (u *User) GetTodosByUser() (todos []Todo, err error) {
	//SQL文を指定
	cmd := `select id, content, user_id, created_at from todos
	where user_id = $1`

	//取得処理を実行
	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}

	//タスクの取得処理
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt,
		)

		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

//タスクの更新処理
func (t *Todo) UpdateTodo() error {
	//SQL文を指定
	cmd := `update todos set content = $1, user_id = $2 where id = $3`

	//更新処理を実行
	_, err = Db.Exec(cmd, t.Content, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

//タスクの削除処理
func (t *Todo) DeleteTodo() error {
	//SQL文を指定
	cmd := `delete from todos where id = $1`

	//削除処理を実行
	_, err = Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

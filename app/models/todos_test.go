package models

//参考：https://pkg.go.dev/github.com/DATA-DOG/go-sqlmock#section-readme
import (
	"database/sql"
	"database/sql/driver"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/assert"
)

//テストに使用するユーザーを定義
var testUser = &User{
	ID:       9999999,
	Name:     "testuser",
	Email:    "tests@exqmaple.com",
	PassWord: "pass",
}

//テストに使用するタスクを定義
var testTodo = &Todo{
	ID:        1000000,
	Content:   "testTask",
	UserID:    9999999,
	CreatedAt: time.Now(),
}

//mockを作成する
func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

//登録処理のテスト
func TestTodoCreate(t *testing.T) {
	cmd := "insert into todos (content,user_id,created_at) values (?,?,?)"

	//dbモックを作成
	Db, mock := NewMock()
	defer Db.Close()

	//SQL実行の期待値を定義
	mock.ExpectExec(regexp.QuoteMeta(cmd)).
		WithArgs(testTodo.Content, testUser.ID, AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))

	//作成処理を実行
	//メソッド内のSQLでtime.Now()を使用している場合はクエリ自体のテストのみ行う
	_, err = Db.Exec(cmd, testTodo.Content, testUser.ID, time.Now())
	if err != nil {
		t.Error(err.Error())
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

//IDに応じた取得処理のテスト
func TestTodoGet(t *testing.T) {
	cmd := "select id, content, user_id, created_at from todos where id = ?"

	//dbモックを作成
	Db, mock := NewMock()
	defer Db.Close()

	//取得の期待値を定義
	mock.ExpectQuery(regexp.QuoteMeta(cmd)).
		WithArgs(testTodo.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "content", "user_id", "created_at"}).AddRow(testTodo.ID, testTodo.Content, testUser.ID, testTodo.CreatedAt))

	//取得の実行
	todo, err := GetTodo(Db, testTodo.ID)
	assert.Nil(t, err)
	assert.Nil(t, deep.Equal(todo.Content, testTodo.Content))

}

//全取得処理のテスト
func TestTodosGet(t *testing.T) {
	cmd := `select id, content, user_id, created_at from todos`

	//dbモックを作成
	Db, mock := NewMock()
	defer Db.Close()

	//取得の期待値を定義
	mock.ExpectQuery(regexp.QuoteMeta(cmd)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "content", "user_id", "created_at"}).AddRow(testTodo.ID, testTodo.Content, testUser.ID, testTodo.CreatedAt))

	//取得の実行
	todo, err := GetTodos(Db)
	assert.Nil(t, err)
	assert.Nil(t, deep.Equal(todo[0].Content, testTodo.Content))
}

//ユーザーに応じた取得処理のテスト
func TestTodoGetByUser(t *testing.T) {
	cmd := `select id, content, user_id, created_at from todos
	where user_id = ?`

	//dbモックを作成
	Db, mock := NewMock()
	defer Db.Close()

	//取得の期待値を定義
	mock.ExpectQuery(regexp.QuoteMeta(cmd)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "content", "user_id", "created_at"}).AddRow(testTodo.ID, testTodo.Content, testUser.ID, testTodo.CreatedAt))

	//取得の実行
	todo, err := testUser.GetTodosByUser(Db)
	assert.Nil(t, err)
	assert.Nil(t, deep.Equal(todo[0].Content, testTodo.Content))
}

//更新処理のテスト
func TestTodoUpdate(t *testing.T) {
	cmd := "update todos set content = ?, user_id = ? where id = ?"
	cmdSelect := `select id, content, user_id, created_at from todos where id = ?`
	afterContent := "Updated Task"

	//dbモックを作成
	Db, mock := NewMock()
	defer Db.Close()

	//SQL実行の期待値を定義
	mock.ExpectExec(regexp.QuoteMeta(cmd)).
		WithArgs(afterContent, testUser.ID, testTodo.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	//更新処理を実行
	testTodoUpdate := testTodo
	testTodoUpdate.Content = afterContent
	err := testTodoUpdate.UpdateTodo(Db)
	if err != nil {
		t.Error(err.Error())
	}

	//更新結果の取得
	mock.ExpectQuery(regexp.QuoteMeta(cmdSelect)).
		WithArgs(testTodo.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "content", "user_id", "created_at"}).AddRow(testTodo.ID, afterContent, testTodo.UserID, testTodo.CreatedAt))

	//取得の実行
	todo, err := GetTodo(Db, testTodo.ID)
	assert.Nil(t, err)
	assert.Nil(t, deep.Equal(todo.Content, afterContent))

}

//削除処理のテスト
func TestTodoDelete(t *testing.T) {
	cmd := "delete from todos where id = ?"

	//dbモックを作成
	Db, mock := NewMock()
	defer Db.Close()

	//SQL実行の期待値を定義
	mock.ExpectExec(regexp.QuoteMeta(cmd)).
		WithArgs(testTodo.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	//削除処理の実行
	err := testTodo.DeleteTodo(Db)
	assert.Nil(t, err)
}

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

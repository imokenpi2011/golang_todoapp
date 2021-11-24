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
	ID:        9999999,
	UUID:      "testUUID",
	Name:      "test",
	Email:     "test@exqmaple.com",
	PassWord:  "pass",
	CreatedAt: time.Now(),
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
func TestUserCreate(t *testing.T) {
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

	//dbモックを作成
	Db, mock := NewMock()
	defer Db.Close()

	//SQL実行の期待値を定義
	mock.ExpectExec(regexp.QuoteMeta(cmd)).
		WithArgs(testUser.UUID, testUser.Name, testUser.Email, testUser.PassWord, AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))

	//作成処理を実行
	//メソッド内のSQLでtime.Now()を使用している場合はクエリ自体のテストのみ行う
	_, err = Db.Exec(cmd, testUser.UUID, testUser.Name, testUser.Email, testUser.PassWord, testUser.CreatedAt)
	if err != nil {
		t.Error(err.Error())
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

//IDに応じた取得処理のテスト
func TestUserGet(t *testing.T) {
	cmd := "select id, content, user_id, created_at from todos where id = ?"

	//dbモックを作成
	Db, mock := NewMock()
	defer Db.Close()

	//仮ユーザーを作成
	err := testUser.CreateUser(Db)
	if err != nil {
		t.Error(err.Error())
	}

	//取得の期待値を定義
	mock.ExpectQuery(regexp.QuoteMeta(cmd)).
		WithArgs(testTodo.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "uuid", "name", "email", "password", "created_at"}).
			AddRow(testUser.ID, testUser.UUID, testUser.Name, testUser.Email, testUser.PassWord, testUser.CreatedAt))

	//取得の実行
	user, err := GetUser(Db, 1)
	assert.Nil(t, err)
	assert.Nil(t, deep.Equal(user.Name, testUser.Name))
	assert.Nil(t, deep.Equal(user.Email, testUser.Email))
	assert.Nil(t, deep.Equal(user.PassWord, Encrypt(testUser.PassWord)))
}

//ユーザーに応じた取得処理のテスト
func TestUserGetByEmail(t *testing.T) {
	cmd := `select id, uuid, name, email, password, created_at
	from users where email = ?`

	//dbモックを作成
	Db, mock := NewMock()
	defer Db.Close()

	//仮ユーザーを作成
	err := testUser.CreateUser(Db)
	if err != nil {
		t.Error(err.Error())
	}

	//作成したユーザーを取得する
	createdUser, err := GetUser(Db, 1)
	if err != nil {
		t.Error(err.Error())
	}

	//取得の期待値を定義
	mock.ExpectQuery(regexp.QuoteMeta(cmd)).
		WithArgs(testUser.Email).
		WillReturnRows(sqlmock.NewRows([]string{"id", "uuid", "name", "email", "password", "created_at"}).
			AddRow(createdUser.ID, createdUser.UUID, createdUser.Name, createdUser.Email, createdUser.PassWord, createdUser.CreatedAt))

	//取得の実行
	user, err := GetUserByEmail(Db, createdUser.Email)
	assert.Nil(t, err)
	assert.Nil(t, deep.Equal(user, createdUser))
}

//更新処理のテスト
func TestUserUpdate(t *testing.T) {
	cmd := `update users set name = ?, email = ? where id = ?`
	cmdSelect := `select id, uuid, name, email, password, created_at from users where id = ?`
	afterUser := &User{
		ID:        testUser.ID,
		UUID:      testUser.UUID,
		Name:      "afterUser",
		Email:     "updated@exqmaple.com",
		PassWord:  testUser.PassWord,
		CreatedAt: testUser.CreatedAt,
	}

	//dbモックを作成
	Db, mock := NewMock()
	defer Db.Close()

	//SQL実行の期待値を定義
	mock.ExpectExec(regexp.QuoteMeta(cmd)).
		WithArgs(afterUser.Name, afterUser.Email, afterUser.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	//更新処理を実行
	err := afterUser.UpdateUser(Db)
	if err != nil {
		t.Error(err.Error())
	}

	//更新結果の取得
	mock.ExpectQuery(regexp.QuoteMeta(cmdSelect)).
		WithArgs(afterUser.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "uuid", "name", "email", "password", "created_at"}).
			AddRow(afterUser.ID, afterUser.UUID, afterUser.Name, afterUser.Email, afterUser.PassWord, afterUser.CreatedAt))

	//取得の実行
	user, err := GetUser(Db, afterUser.ID)
	// user := User{}
	// err = Db.QueryRow(cmdSelect, afterUser.ID).Scan(
	// 	&user.ID,
	// 	&user.UUID,
	// 	&user.Name,
	// 	&user.Email,
	// 	&user.PassWord,
	// 	&user.CreatedAt,
	// )
	assert.Nil(t, err)
	assert.Nil(t, deep.Equal(user.Name, afterUser.Name))
	assert.Nil(t, deep.Equal(user.Email, afterUser.Email))

}

//削除処理のテスト
func TestUserDelete(t *testing.T) {
	cmd := `delete from users where id = ?`

	//dbモックを作成
	Db, mock := NewMock()
	defer Db.Close()

	//SQL実行の期待値を定義
	mock.ExpectExec(regexp.QuoteMeta(cmd)).
		WithArgs(testUser.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	//削除処理の実行
	err := testUser.DeleteUser(Db)
	assert.Nil(t, err)
}

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
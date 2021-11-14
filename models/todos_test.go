package models

//参考：https://uzimihsr.github.io/post/2021-04-30-golang-test-with-go-sqlmock/
import (
	"errors"
	"golang_todoapp/clock"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

//タスク登録に使用するユーザーを定義
var u = &User{
	ID:       9999999,
	Name:     "testuser",
	Email:    "tests@exqmaple.com",
	PassWord: "pass",
}

var taskName string = "testTask"

//登録処理のテスト
func TestTodoCreate(t *testing.T) {
	cmd := "insert into todos \\(content,user_id,created_at\\) values \\(?,?,?\\)"

	clock.Set((time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)))

	t.Run(
		"Create_正常系",
		func(t *testing.T) {
			//dbモックを作成
			Db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer Db.Close()

			//SQL実行の期待値を定義
			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta(cmd)).
				WithArgs(taskName, u.ID, clock.Now()).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			//作成処理を実行
			err = u.CreateTodo(Db, taskName)
			if err != nil {
				t.Error(err.Error())
			}
		},
	)

	t.Run(
		"Create_異常系",
		func(t *testing.T) {
			//dbモックを作成
			Db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer Db.Close()

			//SQL実行時にエラーを起こすよう定義
			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta(cmd)).
				WithArgs(taskName, u.ID, clock.Now()).
				WillReturnResult(sqlmock.NewErrorResult(errors.New("ERROR!!!"))).
				WillReturnError(errors.New("INSERT FAILED!!!"))
			mock.ExpectRollback()

			//作成処理を実行
			err = u.CreateTodo(Db, taskName)

			//正常終了した場合はエラー
			if err == nil {
				t.Error("An error should have occurred.")
			}
		},
	)
}

// //登録するタスクを定義
// var taskList = []string{"task01", "task02"}

// //タスク更新時の定数
// var updateContent = "updatetask"

// //todo関連のcrudテストを行う
// func TestTodos(t *testing.T) {

// 	//作成処理テスト
// 	for _, task := range taskList {
// 		err = u.CreateTodo(task)
// 		if err != nil {
// 			t.Errorf("Failed create todo. err = %v", err)
// 		}
// 	}

// 	//全タスクの取得テスト
// 	allTodos, err := GetTodos()
// 	if err != nil {
// 		t.Errorf("Failed get all todos. err = %v", err)
// 	} else if len(allTodos) == 0 {
// 		t.Errorf("Failed to get registed todos.")
// 	}

// 	//作成したタスクの確認処理
// 	todos, err := u.GetTodosByUser()
// 	if err != nil {
// 		t.Errorf("Failed get todos by userID. err = %v", err)
// 	}
// 	for i, todo := range todos {
// 		t.Logf("ID:%v / content:%v / userID:%v", strconv.Itoa(todo.ID), todo.Content, strconv.Itoa(todo.UserID))
// 		//削除処理テスト
// 		defer func() {
// 			//最後に今回作成したタスク情報を削除する
// 			err = todo.DeleteTodo()
// 			if err != nil {
// 				t.Errorf("Failed delete todo. err = %v", err)
// 			}
// 		}()

// 		//単一タスク取得テスト
// 		todo, err := GetTodo(todo.ID)
// 		if err != nil {
// 			t.Errorf("Failed get todo. err = %v", err)
// 		} else if taskList[i] != todo.Content {
// 			//想定通りの値が登録されているか確認
// 			t.Errorf("There are discrepancies in resisted values.  expected =%v  / registed = %v", taskList[i], todo.Content)
// 		}
// 		fmt.Println(todo.Content)

// 	}

// 	//タスクの更新処理テスト
// 	todos[0].Content = updateContent
// 	err = todos[0].UpdateTodo()
// 	if err != nil {
// 		t.Errorf("Failed update todo. err = %v", err)
// 	} else {
// 		//文言が更新されているか確認
// 		updatedTodo, _ := GetTodo(todos[0].ID)
// 		if updatedTodo.Content != updateContent {
// 			t.Logf("There are discrepancies in updated values.  expected =%v  / registed = %v", updateContent, updatedTodo.Content)
// 		}
// 	}
//}

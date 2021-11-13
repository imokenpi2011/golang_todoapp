package main

import (
	"fmt"
	"golang_todoapp/models"
)

func main() {
	//modelsパッケージで指定したDbの読み込み値を出力
	fmt.Println(models.Db)
}

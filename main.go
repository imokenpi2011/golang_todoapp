package main

import (
	"fmt"
	"golang_todoapp/config"
	"log"
)

func main() {
	//設定値を出力
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)

	//ログを出力
	log.Println("test")
}

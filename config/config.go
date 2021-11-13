package config

import (
	"golang_todoapp/utils"
	"log"

	"gopkg.in/go-ini/ini.v1"
)

// 設定一覧を指定
type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
	Static    string
}

//Configインスタンスを定義
var Config ConfigList

//初期処理
func init() {
	//設定の読み込み処理
	LoadConfig()
	//ログの読み込みを設定
	utils.LoggingSettings(Config.LogFile)
}

//設定を読み込む
func LoadConfig() {
	// 設定を記載したiniファイルを読み込む
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}

	//読み込んだ値を設定する
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		DbName:    cfg.Section("db").Key("name").String(),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
		Static:    cfg.Section("web").Key("static").String(),
	}
}

package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	//ログファイルを読み込む。ファイルが無い場合は作成する
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	//出力先の指定
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	//ログフォーマット指定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//出力先指定
	log.SetOutput(multiLogFile)
}

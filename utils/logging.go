package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)   // ログを標準出力とlogfileへの出力に指定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) // ログフォーマット -> 2023/12/21 21:03:30 main.go:15: test
	log.SetOutput(multiLogFile)                          // ログ出力先
}

package controllers

import (
	"net/http"
	"todo_app/config"
)

func StartMainServer() error {

	// URL登録
	http.HandleFunc("/", top)

	// 第二引数はnilを入れてデフォルトのマルチプレクサを使う
	// アクセスしたことがないURLへアクセスした場合、PageNotFoundを返す仕組み
	return http.ListenAndServe(":"+config.Config.Port, nil)
}

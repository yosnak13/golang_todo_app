package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"todo_app/app/models"
	"todo_app/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)

}

func getCookie(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("invalid getCookie")
		}
	}
	return sess, err
}

func StartMainServer() error {

	// Static読み込み。"/static/"ディレクトリはないためStripPrefixで取り除く
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// URL登録
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)

	// 第二引数はnilを入れてデフォルトのマルチプレクサを使う
	// アクセスしたことがないURLへアクセスした場合、PageNotFoundを返す仕組み
	return http.ListenAndServe(":"+config.Config.Port, nil)
}

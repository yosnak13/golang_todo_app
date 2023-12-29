package controllers

import (
	"log"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	_, err := getCookie(w, r)
	if err != nil {
		generateHTML(w, "Hello", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	sess, err := getCookie(w, r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/top", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		todos, _ := user.ListTodosByUser()
		user.Todos = todos
		// 第二引数はフロントに渡したいデータを渡す
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}

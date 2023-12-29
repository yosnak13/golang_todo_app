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
	_, err := getCookie(w, r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/top", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "index")
	}
}

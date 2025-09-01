package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Page struct {
	Title string
	Body  string
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "HELLO\n")
}

func index(w http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := Page{
		Title: "My first blog post",
		Body:  "Hello, world!",
	}

	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8088", nil)
}

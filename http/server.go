package http

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func GetIndexPage(w http.ResponseWriter, req *http.Request) {
	fmt.Println("index")
	tmpl, err := template.ParseFiles("http/templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	posts := GetPosts()
	tmpl.Execute(w, posts)
}

func GetPost(w http.ResponseWriter, req *http.Request) {
	fmt.Println("GET post", req.PathValue("id"))
	tmpl, err := template.ParseFiles("http/templates/post.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id := req.PathValue("id")
	post, err := GetPostById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, post)
}

func SubmitPost(w http.ResponseWriter, req *http.Request) {
	fmt.Println("POST post")
	tmpl, err := template.ParseFiles("http/templates/admin/submit-post.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if req.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	post := Post{
		Title:   req.FormValue("title"),
		Summary: req.FormValue("summary"),
		Body:    req.FormValue("body"),
		ID:      strconv.Itoa(len(posts)),
	}
	AddPost(post)
	tmpl.Execute(w, struct{ Success bool }{true})
}

func EditPost(w http.ResponseWriter, req *http.Request) {
	fmt.Println("UPDATE post")
	tmpl, err := template.ParseFiles("http/templates/admin/submit-post.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id := req.PathValue("id")
	ogPost, err := GetPostById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if req.Method != http.MethodPost {
		tmpl.Execute(w, ogPost)
		return
	}

	post := Post{
		Title:   req.FormValue("title"),
		Summary: req.FormValue("summary"),
		Body:    req.FormValue("body"),
		ID:      strconv.Itoa(len(posts)),
	}
	UpdatePost(post)
	tmpl.Execute(w, struct{ Success bool }{true})
}

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "HELLO, WORLD\n")
}

package http

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func check(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetIndexPage(w http.ResponseWriter, req *http.Request) {
	fmt.Println("GET index")
	tmpl, err := template.ParseFiles("http/templates/index.html")
	check(err, w)

	posts, err := GetAllPosts()
	check(err, w)

	tmpl.Execute(w, posts)
}

func GetPost(w http.ResponseWriter, req *http.Request) {
	fmt.Println("GET post", req.PathValue("id"))
	tmpl, err := template.ParseFiles("http/templates/post.html")
	check(err, w)

	idArg := req.PathValue("id")
	id, err := strconv.Atoi(idArg)
	check(err, w)
	post, err := GetPostById(id)
	check(err, w)

	tmpl.Execute(w, post)
}

func SubmitPost(w http.ResponseWriter, req *http.Request) {
	fmt.Println("POST post")
	tmpl, err := template.ParseFiles("http/templates/admin/submit-post.html")
	check(err, w)

	if req.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	AddPost(req.FormValue("title"), req.FormValue("summary"), req.FormValue("body"))
	tmpl.Execute(w, struct{ Success bool }{true})
}

func EditPost(w http.ResponseWriter, req *http.Request) {
	fmt.Println("UPDATE post")
	tmpl, err := template.ParseFiles("http/templates/admin/submit-post.html")
	check(err, w)

	idArg := req.PathValue("id")
	id, err := strconv.Atoi(idArg)
	check(err, w)
	ogPost, err := GetPostById(id)
	check(err, w)

	if req.Method != http.MethodPost {
		tmpl.Execute(w, ogPost)
		return
	}

	post := Post{
		Title:   req.FormValue("title"),
		Summary: req.FormValue("summary"),
		Body:    req.FormValue("body"),
		ID:      id,
	}
	UpdatePost(post)
	tmpl.Execute(w, struct{ Success bool }{true})
}

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "HELLO, WORLD\n")
}

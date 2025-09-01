package http

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Body    string `json:"body"`
}

type PostData struct {
	Posts []Post `json:"posts"`
}

var posts []Post

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "HELLO\n")
}

func readPostsFromFile() {
	if len(posts) > 0 {
		return
	}

	dataFile := "./posts.json"
	dat, err := os.ReadFile(dataFile)
	if err != nil {
		panic(err)
	}

	var jsonData PostData
	if err := json.Unmarshal(dat, &jsonData); err != nil {
		panic(err)
	}
	posts = append(posts, jsonData.Posts...)
}

func GetIndex(w http.ResponseWriter, req *http.Request) {
	fmt.Println("index")
	tmpl, err := template.ParseFiles("http/templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	readPostsFromFile()
	tmpl.Execute(w, posts)
}

func GetPost(w http.ResponseWriter, req *http.Request) {
	fmt.Println("post", req.PathValue("id"))
	tmpl, err := template.ParseFiles("http/templates/post.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	readPostsFromFile()
	id := req.PathValue("id")
	for _, post := range posts {
		if post.ID == id {
			tmpl.Execute(w, post)
			return
		}
	}
}

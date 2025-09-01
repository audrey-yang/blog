package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "HELLO\n")
}

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

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Println("index")
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	readPostsFromFile()
	tmpl.Execute(w, posts)
}

func post(w http.ResponseWriter, req *http.Request) {
	fmt.Println("post", req.PathValue("id"))
	tmpl, err := template.ParseFiles("templates/post.html")
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

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/post/{id}", post)
	mux.HandleFunc("/hello", hello)

	fmt.Println("Listening...")
	if err := http.ListenAndServe(":8090", mux); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}

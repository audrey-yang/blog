package http

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
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

var dataFile string = "./posts.json"
var posts []Post

func readPostsFromFile() {
	if len(posts) > 0 {
		return
	}

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

func writePostsToFile() {
	dataJson, jsonErr := json.MarshalIndent(posts, "", "  ")
	if jsonErr != nil {
		panic(jsonErr)
	}
	writeErr := os.WriteFile(dataFile, dataJson, 0644)
	if writeErr != nil {
		panic(writeErr)
	}
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
	fmt.Println("GET post", req.PathValue("id"))
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

func SubmitPost(w http.ResponseWriter, req *http.Request) {
	fmt.Println("POST post")
	tmpl, err := template.ParseFiles("http/templates/submit-post.html")
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
	posts = append(posts, post)

	writePostsToFile()
	tmpl.Execute(w, struct{ Success bool }{true})
}

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "HELLO, WORLD\n")
}

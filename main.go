package main

import (
	server "blog/http"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/post/{id}", server.GetPost)
	mux.HandleFunc("/editor/post", server.SubmitPost)
	mux.HandleFunc("/hello", server.Hello)

	mux.HandleFunc("/", server.GetIndex)

	fmt.Println("Listening...")
	if err := http.ListenAndServe(":8090", mux); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}

package main

import (
	server "blog/http"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", server.GetIndex)
	mux.HandleFunc("/post/{id}", server.GetPost)
	mux.HandleFunc("/hello", server.Hello)

	fmt.Println("Listening...")
	if err := http.ListenAndServe(":8090", mux); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}

package http

import (
	"encoding/json"
	"errors"
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
	Posts []Post
}

var dataFile string = "./posts.json"
var posts []Post

func GetPosts() []Post {
	if len(posts) > 0 {
		return posts
	}
	return readPostsFromFile()
}

func GetPostById(id string) (Post, error) {
	posts := GetPosts()
	for _, post := range posts {
		if post.ID == id {
			return post, nil
		}
	}
	return Post{}, errors.New("could not find Post")
}

func AddPost(newPost Post) {
	posts = append(posts, newPost)
	writePostsToFile()
}

func UpdatePost(updatedPost Post) {
	posts := GetPosts()
	ind, err := strconv.Atoi(updatedPost.ID)
	if err != nil {
		return
	}
	posts[ind-1] = updatedPost
	writePostsToFile()
}

func readPostsFromFile() []Post {
	dat, err := os.ReadFile(dataFile)
	if err != nil {
		panic(err)
	}

	var jsonData PostData
	if err := json.Unmarshal(dat, &jsonData); err != nil {
		panic(err)
	}
	posts = append(posts, jsonData.Posts...)
	return posts
}

func writePostsToFile() {
	dataJson, jsonErr := json.MarshalIndent(PostData{posts}, "", "  ")
	if jsonErr != nil {
		panic(jsonErr)
	}
	writeErr := os.WriteFile(dataFile, dataJson, 0644)
	if writeErr != nil {
		panic(writeErr)
	}
}

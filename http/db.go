package http

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn
var ctx = context.Background()

func InitDb() {
	var err error
	conn, err = pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	if err := conn.Ping(ctx); err != nil {
		log.Fatal("Unable to ping database:", err)
	}

	fmt.Println("Connected to PostgreSQL database!")
}

type Post struct {
	ID      int
	Title   string
	Summary string
	Body    string
}

func GetAllPosts() ([]Post, error) {
	sql := `
        SELECT id, title, summary, body
        FROM blog_posts
		ORDER BY updated_at DESC
    `
	rows, err := conn.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("error querying Posts: %w", err)
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Summary,
			&post.Body,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning Post row: %w", err)
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func GetPostById(id int) (Post, error) {
	sql := `
        SELECT id, title, summary, body 
        FROM blog_posts
        WHERE id = $1
    `
	var post Post
	err := conn.QueryRow(ctx, sql, id).Scan(
		&post.ID,
		&post.Title,
		&post.Summary,
		&post.Body,
	)
	if err != nil {
		return Post{}, fmt.Errorf("error querying Post: %w", err)
	}

	return post, nil
}

func AddPost(
	title string,
	summary string,
	body string,
) error {
	sql := `
        INSERT INTO blog_posts (title, summary, body)
        VALUES ($1, $2, $3)
        RETURNING id
    `

	var id int
	err := conn.QueryRow(ctx, sql, title, summary, body).Scan(&id)
	if err != nil {
		return fmt.Errorf("error creating Post: %w", err)
	}

	return nil
}

func UpdatePost(updatedPost Post) error {
	sql := `
        UPDATE blog_posts
        SET title = $2, summary = $3, body = $4
        WHERE id = $1
    `

	commandTag, err := conn.Exec(ctx, sql, updatedPost.ID, updatedPost.Title, updatedPost.Summary, updatedPost.Body)
	if err != nil {
		return fmt.Errorf("error completing Post: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("no Post found with id %d", updatedPost.ID)
	}

	return nil
}

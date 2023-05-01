package posts

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

var db *sql.DB

const queryString string = `SELECT posts.id, posts.body, posts.likes, posts.author, users.username 
FROM posts
LEFT JOIN users 
ON posts.author = users.id`

type Post struct {
	ID       int    `json:"id"`
	Body     string `json:"body"`
	Likes    int    `json:"likes"`
	Author   int    `json:"author"`
	Username string `json:"username"`
}

type PostPayload struct {
	Body   string `json:"body"`
	Author int    `json:"author"`
}

func CreatePost(c *fiber.Ctx) error {
	data := PostPayload{}
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	statement := `INSERT INTO posts (body, author) VALUES ($1, $2)`
	_, err = db.Exec(statement, data.Body, data.Author)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(201).SendString("Post created successfully")
}

func GetAllPosts(c *fiber.Ctx) error {
	rows, err := db.Query(queryString)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	posts := []Post{}
	for rows.Next() {
		post := Post{}
		err := rows.Scan(&post.ID, &post.Body, &post.Likes, &post.Author, &post.Username)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		posts = append(posts, post)
	}

	return c.JSON(posts)
}

func GetPost(c *fiber.Ctx) error {
	statement := queryString + `
	WHERE posts.id=(?)`
	row := db.QueryRow(statement, c.Query("id"))

	post := Post{}
	err := row.Scan(&post.ID, &post.Body, &post.Likes, &post.Author, &post.Username)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(post)
}

func LikePost(c *fiber.Ctx) error {
	data := struct {
		ID int `json:"id"`
	}{}

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	statement := `
	UPDATE posts
	SET likes = likes + 1
	WHERE posts.id=(?)`
	_, err = db.Exec(statement, data.ID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendString("Post liked successfully")
}

func RegisterRoutes(router fiber.Router, database *sql.DB) {
	db = database
	posts := router.Group("posts")
	posts.Get("/", GetAllPosts)
	posts.Get("/one", GetPost)
	posts.Post("/", CreatePost)
	posts.Patch("/like", LikePost)
}

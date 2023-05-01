package users

import (
	"database/sql"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID   int    `json:"id"`
	Username string `json:"username"`
}

type UserPayload struct {
	Username string `json:"username"`
}

var db *sql.DB

func CreateUser(c *fiber.Ctx) error {
	data := UserPayload{}
	err := json.Unmarshal(c.Body(), &data)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	statement := `INSERT INTO users (username) VALUES ($1)`
	_, err = db.Exec(statement, data.Username)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(201).SendString("Post created successfully")
}

func GetAllUsers(c *fiber.Ctx) error {
	statement := `SELECT * FROM users`
	rows, err := db.Query(statement)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	users := []User{}
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Username)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		users = append(users, user)
	}

	return c.JSON(users)
}

func RegisterRoutes(router fiber.Router, database *sql.DB) {
	db = database
	users := router.Group("users")
	users.Get("/", GetAllUsers)
	users.Post("/", CreateUser)
}

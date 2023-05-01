package main

import (
	"database/sql"
	"log"

	"github.com/raxracks/chirpr/frontend/pages"
	"github.com/raxracks/chirpr/posts"
	"github.com/raxracks/chirpr/users"
	"github.com/gofiber/fiber/v2"
	_ "modernc.org/sqlite"
)

func main() {
	app := fiber.New()

	// database setup
	db, err := sql.Open("sqlite", "data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		body TEXT NOT NULL,
		likes INTEGER DEFAULT 0,
		author INTEGER,
		FOREIGN KEY (author) REFERENCES users(id)
	);
	VACUUM;`)
	if err != nil {
		log.Fatal(err)
	}

	// routes
	pages.RegisterRoutes(app)

	api := app.Group("api")
	v1 := api.Group("v1")

	users.RegisterRoutes(v1, db)
	posts.RegisterRoutes(v1, db)

	app.Listen(":3000")
}

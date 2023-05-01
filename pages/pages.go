package pages

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	app.Static("assets", "assets")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("pages/index.html")
	})
}

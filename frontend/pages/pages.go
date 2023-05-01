package pages

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	app.Static("assets", "frontend/assets")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("frontend/pages/index.html")
	})
}

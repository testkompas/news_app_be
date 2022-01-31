package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New(fiber.Config{})

	app.Post("/login", func(c *fiber.Ctx) error {
		response := fiber.Map{
			"result": "success",
			"error":  nil,
		}

		return c.JSON(response)
	})

	app.Listen(":8100")
}

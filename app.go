package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{})

	user := os.Getenv("MYSQL_USER")
	if user == "" {
		panic("MYSQL_USER not exists")
	}
	password := os.Getenv("MYSQL_PASS")
	if password == "" {
		panic("MYSQL_PASS not exists")
	}
	address := os.Getenv("MYSQL_ADDRESS")
	if address == "" {
		panic("MYSQL_ADDRESS not exists")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		panic("DB_NAME not exists")
	}

	fmt.Printf("%s:%s@tcp(%s)/%s", user, password, address, dbName)

	app.Post("/login", func(c *fiber.Ctx) error {
		response := fiber.Map{
			"result": "success",
			"error":  nil,
		}

		return c.JSON(response)
	})

	app.Listen(":8100")
}

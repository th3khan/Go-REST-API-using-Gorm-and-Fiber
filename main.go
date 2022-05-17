package main

import "github.com/gofiber/fiber/v2"

func helloWorld(c *fiber.Ctx) error {
	c.SendString("Hello, World!")
	return nil
}

func main() {
	app := fiber.New()

	app.Get("/", helloWorld)

	app.Listen(":3000")
}

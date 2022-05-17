package book

import "github.com/gofiber/fiber/v2"

func GetBooks(c *fiber.Ctx) error {
	return c.SendString("Get All Books")
}

func GetBook(c *fiber.Ctx) error {
	return c.SendString("Get a Single Book")
}

func NewBook(c *fiber.Ctx) error {
	return c.SendString("Create a New Book")
}

func UpdateBook(c *fiber.Ctx) error {
	return c.SendString("Update a Book")
}

func DeleteBook(c *fiber.Ctx) error {
	return c.SendString("Delete a Book")
}

package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/Go-REST-API-using-Gorm-and-Fiber/book"
	"github.com/th3khan/Go-REST-API-using-Gorm-and-Fiber/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func helloWorld(c *fiber.Ctx) error {
	c.SendString("Hello, World!")
	return nil
}

func setupRoutes(app *fiber.App) {
	// books
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Put("/api/v1/book/:id", book.UpdateBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	// connect to database
	var err error
	database.DbConnection, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connected to database successfully!")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)

	app.Listen(":3000")
}

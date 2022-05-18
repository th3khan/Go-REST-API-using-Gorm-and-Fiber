package book

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/Go-REST-API-using-Gorm-and-Fiber/database"
	"gorm.io/gorm"
)

type Book struct {
	ID        uint           `gorm:"primary_key" json:"id"`
	Title     string         `json:"title"`
	Author    string         `json:"author"`
	Rating    int            `json:"rating"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeleteAt  gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at" gorm:"index"`
}

func GetBooks(c *fiber.Ctx) error {
	db := database.DbConnection
	var books []Book
	db.Find(&books)
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DbConnection
	var book Book
	db.First(&book, id)

	if book.ID == 0 {
		return c.SendStatus(404)
	}
	return c.JSON(book)
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

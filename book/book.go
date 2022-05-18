package book

import (
	"errors"
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
	DeleteAt  gorm.DeletedAt `gorm:"column:deleted_at" gorm:"index"`
}

type BookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) error {
	db := database.DbConnection
	var books []Book
	db.Find(&books)
	return c.JSON(books)
}

func ValidateBookExists(c *fiber.Ctx, db *gorm.DB) (Book, error, int) {
	id, err := c.ParamsInt("id")
	if err != nil {
		return Book{}, errors.New("Invalid request"), fiber.StatusBadRequest
	}
	var book Book
	db.First(&book, id)
	if book.ID == 0 {
		return Book{}, errors.New("Book not found"), fiber.StatusNotFound
	}
	return book, nil, fiber.StatusOK
}

func GetBook(c *fiber.Ctx) error {
	db := database.DbConnection
	book, err, status := ValidateBookExists(c, db)

	if err != nil {
		c.JSON(fiber.Map{
			"message": err.Error(),
		})
		return c.SendStatus(status)
	}
	return c.JSON(book)
}

func NewBook(c *fiber.Ctx) error {
	var book BookRequest
	if err := c.BodyParser(&book); err != nil {
		c.JSON(fiber.Map{
			"message": "Invalid request",
		})
		return c.SendStatus(fiber.StatusBadRequest)
	}
	db := database.DbConnection

	var newBook Book
	newBook.Title = book.Title
	newBook.Author = book.Author
	newBook.Rating = book.Rating
	db.Create(&newBook)

	c.JSON(newBook)
	return c.SendStatus(fiber.StatusCreated)
}

func UpdateBook(c *fiber.Ctx) error {
	db := database.DbConnection
	book, err, status := ValidateBookExists(c, db)
	if err != nil {
		c.JSON(fiber.Map{
			"message": err.Error(),
		})
		return c.SendStatus(status)
	}
	var bookRequest BookRequest
	if err := c.BodyParser(&bookRequest); err != nil {
		c.JSON(fiber.Map{
			"message": "Invalid request",
		})
		return c.SendStatus(fiber.StatusBadRequest)
	}
	book.Title = bookRequest.Title
	book.Author = bookRequest.Author
	book.Rating = bookRequest.Rating
	db.Save(&book)
	c.JSON(book)
	return c.SendStatus(fiber.StatusOK)
}

func DeleteBook(c *fiber.Ctx) error {
	db := database.DbConnection
	book, err, status := ValidateBookExists(c, db)
	if err != nil {
		c.JSON(fiber.Map{
			"message": err.Error(),
		})
		return c.SendStatus(status)
	}
	db.Delete(&book)
	c.JSON(fiber.Map{
		"message": "Book deleted",
	})
	return c.SendStatus(fiber.StatusOK)
}

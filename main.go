package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"main.go/database"
	"main.go/models"
)

func main() {
	database.Connect()
	defer database.CloseDatabase(database.Connect())
	database.Migrations()
	app := fiber.New()

	app.Get("/gets", func(c *fiber.Ctx) error {
		db := database.Connect()
		defer database.CloseDatabase(db)

		var book []models.Book
		db.Find(&book)
		return c.JSON(book)
	})

	app.Post("/post", func(c *fiber.Ctx) error {
		db := database.Connect()
		defer database.CloseDatabase(db)

		book := new(models.Book)
		if err := c.BodyParser(book); err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "Error",
				"data":    err,
			})
		}
		db.Create(&book)
		return c.JSON(book)
	})

	app.Get("/get/:id", func(c *fiber.Ctx) error {
		db := database.Connect()
		defer database.CloseDatabase(db)

		id := c.Params("id")
		var book models.Book
		db.Find(&book, id)
		return c.JSON(book)
	})

	app.Delete("/delete/:id", func(c *fiber.Ctx) error {
		db := database.Connect()

		id := c.Params("id")
		var book models.Book
		err := db.Find(&book, "id = ?", id).Error
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "error in delete func",
			})
		}
		db.Delete(&book)
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Data deleted",
		})
	})

	app.Put("/update/:id", func(c *fiber.Ctx) error {
		type UpdateBook struct {
			Title  string `json:"title"`
			Author string `json:"author"`
			Rating int    `json:"rating"`
		}

		db := database.Connect()

		var book models.Book
		id := c.Params("id")
		db.First(&book, "id = ?", id)

		var updateBook UpdateBook
		err := c.BodyParser(&updateBook)
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "Review your inputs",
				"data":    err,
			})
		}

		book.Title = updateBook.Title
		book.Author = updateBook.Author
		book.Rating = updateBook.Rating

		db.Save(&book)

		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "category found",
			"data":    book,
		})
	})

	log.Fatal(app.Listen(":3000"))
}

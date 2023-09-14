package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	db "github.com/nyshnt/go-fiber-postgres/database"
	"github.com/nyshnt/go-fiber-postgres/models"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

func (b *Book) CreateBook(context *fiber.Ctx) error {
	err := context.BodyParser(&b)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"messae": "request failed"})
		return err
	}

	err = db.DB.Create(&b).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create book"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Book had been added"})
	return nil
}

func (b *Book) DeleteBook(context *fiber.Ctx) error {
	bookModel := models.Books{} // initializes bookModel with empty values
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
	}

	err := db.DB.Delete(bookModel, id)
	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete book",
		})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Book deleted successfully",
	})

	return err.Error
}

func (b *Book) GetBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Books{} // Initializes bookModels with a pointer to an empty slice

	err := db.DB.Find(bookModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get books"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "books fetched successfully",
		"data":    bookModels,
	})
	return nil
}

func (b *Book) GetBookByID(context *fiber.Ctx) error {
	id := context.Params("id")

	bookModel := &models.Books{} // bookModel is a pointer to a new instance of models.Books
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := db.DB.Where("id = ?", id).First(bookModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not get the book ",
		})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Book id fetched successfully",
		"data":    bookModel,
	})
	return nil
}

package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nyshnt/go-fiber-postgres/controllers"
)

func SetupRoutes(app *fiber.App) {

	bookHandler := &controllers.Book{}

	api := app.Group("/api")
	api.Post("/create_books", bookHandler.CreateBook)
	api.Delete("/delete_books/:id", bookHandler.DeleteBook)
	api.Get("/get_book/:id", bookHandler.GetBookByID)
	api.Get("/books", bookHandler.GetBooks)
}

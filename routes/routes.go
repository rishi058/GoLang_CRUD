package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rishi058/go-project/methods"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_books", methods.CreateBook)
	api.Delete("delete_book/:id", methods.DeleteBook)
	api.Get("/get_books/:id", methods.GetBookByID)
	api.Get("/books", methods.GetBooks)
}
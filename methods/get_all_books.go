package methods

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	db "github.com/rishi058/go-project/db_instance"
	"github.com/rishi058/go-project/models"
)

func GetBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Books{}  		// slice of books

	err := db.Instance.Find(bookModels).Error      //! getting all books 

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
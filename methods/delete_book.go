package methods

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	db "github.com/rishi058/go-project/db_instance"
	"github.com/rishi058/go-project/models"
)

func DeleteBook(context *fiber.Ctx) error {
	bookModel := models.Books{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := db.Instance.Delete(bookModel, id)   //! deleting a book from the table.

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete book",
		})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book delete successfully",
	})
	return nil
}
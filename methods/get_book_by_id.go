package methods

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	db "github.com/rishi058/go-project/db_instance"
	"github.com/rishi058/go-project/models"
)

func GetBookByID(context *fiber.Ctx) error {

	id := context.Params("id")
	bookModel := &models.Books{}

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := db.Instance.Where("id = ?", id).First(bookModel).Error   //! getting book by id.

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the book"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book id fetched successfully",
		"data":    bookModel,
	})
	return nil
}
package methods

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	db "github.com/rishi058/go-project/db_instance"
)

type Book struct {                   // this data_model is recieved from client side. [for inserting book]
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}


func CreateBook(context *fiber.Ctx) ( error ) {   // suppose this a function of the class.
	book := Book{}

	err := context.BodyParser(&book)  // context is json recieved from client is parsed int book-struct

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	
	err = db.Instance.Create(&book).Error  // !inserting in database (postgres)

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create book"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book has been added"})
	return nil
}
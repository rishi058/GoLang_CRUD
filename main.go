package main // name of this file -> i.e use this name to import functions/data of this file

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"            // used for routing (it's a web framework)
	"github.com/joho/godotenv"               // used to load environment variables
	"github.com/rishi058/go-project/models"  // importing local files
	"github.com/rishi058/go-project/storage" // importing local files
	"gorm.io/gorm"                           // used for connecting postgres
)

type Book struct {                   // this data_model is recieved from client side. [for inserting book]
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type DB_instance struct {    // suppose it is a class [of database]
	data *gorm.DB
}

func (db *DB_instance) CreateBook(context *fiber.Ctx) ( error ) {   // suppose this a function of the class.
	book := Book{}

	err := context.BodyParser(&book)  // context is json recieved from client is parsed int book-struct

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = db.data.Create(&book).Error  // !inserting in database (postgres)

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create book"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book has been added"})
	return nil
}

func (db *DB_instance) DeleteBook(context *fiber.Ctx) error {
	bookModel := models.BooksTable{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := db.data.Delete(bookModel, id)   //! deleting a book from the table.

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

func (db *DB_instance) GetBooks(context *fiber.Ctx) error {
	bookModels := &[]models.BooksTable{}  		// slice of books

	err := db.data.Find(bookModels).Error      //! getting all books 

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

func (db *DB_instance) GetBookByID(context *fiber.Ctx) error {

	id := context.Params("id")
	bookModel := &models.BooksTable{}

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	fmt.Println("the ID is", id)

	err := db.data.Where("id = ?", id).First(bookModel).Error   //! getting book by id.

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

func (db DB_instance) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_books", db.CreateBook)
	api.Delete("delete_book/:id", db.DeleteBook)
	api.Get("/get_books/:id", db.GetBookByID)
	api.Get("/books", db.GetBooks)
}

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)  // establishing connection with the postgres DB server.

	if err != nil {
		log.Fatal("could not load the database")
	}

	err = models.MigrateData(db)   // creating tables if not exist.

	if err != nil {
		log.Fatal("could not migrate db")
	}

	allData := DB_instance{
		data: db,
	}

	app := fiber.New()
	allData.SetupRoutes(app)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	app.Listen("0.0.0.0:" + port)

}

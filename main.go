package main

import (
	"github.com/gofiber/fiber/v2" // used for routing (it's a web framework)
	db "github.com/rishi058/go-project/db_instance"
	"github.com/rishi058/go-project/routes" // importing local files
)

func main() {

	db.InitializeDB()

	app := fiber.New()

	routes.SetupRoutes(app)

	// port := os.Getenv("PORT")

	// if port == "" {
	// 	port = "3000"
	// }

	// app.Listen("0.0.0.0:" + port)
	app.Listen(":8000")
}

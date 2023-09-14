package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	db "github.com/nyshnt/go-fiber-postgres/database"
	"github.com/nyshnt/go-fiber-postgres/routes"
)

// CreateServer creates a new Fiber instance
func CreateServer() *fiber.App {
	app := fiber.New()

	return app
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.InitDB(); err != nil {
		panic("Failed to initialize the database: " + err.Error())
	}
	app := CreateServer()

	routes.SetupRoutes(app)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
	app.Listen(":4000")
}

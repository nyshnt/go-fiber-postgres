package tests

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetBooks(t *testing.T) {
	// Create a new Fiber app instance (as in your main application)
	app := fiber.New()

	// Define your routes and controllers (similar to your main application)
	app.Get("/api/books")

	// Simulate a GET request to the /api/books endpoint
	req, _ := http.NewRequest("GET", "/api/books", nil)
	resp, _ := app.Test(req)

	// Assert response status code
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// Add more assertions to validate the response body, headers, etc.
}

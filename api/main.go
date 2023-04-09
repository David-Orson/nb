package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	e "nautilus-billing.com/m/v2/error"
	"nautilus-billing.com/m/v2/handler"
	"nautilus-billing.com/m/v2/store"
)

// main is the entry point for the application
func main() {
	port := ":8085"

	// Create a new Fiber instance
	app := fiber.New()

	// Create a new store instance
	s := store.New()

	// Create a new handler instance
	handler.Setup(app, s)

	// Listen on port 8085
	e.Consume(listen(app, port), "Error starting server: ")
}

func listen(app *fiber.App, port string) error {
	fmt.Println("Listening on: " + port)
	return app.Listen(port)
}

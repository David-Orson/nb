package main

import (
	"fmt"
	e "github.com/David-Orson/nb/error"
	"github.com/David-Orson/nb/handler"
	"github.com/David-Orson/nb/store"
	"github.com/gofiber/fiber/v2"
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

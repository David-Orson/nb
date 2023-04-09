package handler

import (
	e "github.com/David-Orson/nb/error"
	"github.com/David-Orson/nb/store/model"
	"github.com/gofiber/fiber/v2"
)

type AccountHandler struct {
	*Handler
}

// Init sets up the routes for the account resource
func (h *AccountHandler) Init() {
	accounts := h.App.Group("/account")

	accounts.Post("/create", h.create)
}

// create handles the creation of a new account
func (h *AccountHandler) create(f *fiber.Ctx) error {
	acc := &model.Account{}

	// Parse the request body into the account model
	if !e.Process(f.BodyParser(acc), "Error parsing account: ") {
		return f.Status(400).SendString("Error parsing account")
	}

	// Create the account
	if !h.Store.Account().Create(acc) {
		return f.Status(500).SendString("Error creating account")
	}

	return f.JSON(acc)
}

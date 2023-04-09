package handler

import (
	"github.com/gofiber/fiber/v2"
	"nautilus-billing.com/m/v2/store"
)

type Handler struct {
	App   *fiber.App
	Store store.Actions
}

// Setup initializes the handler for the application
func Setup(app *fiber.App, s store.Actions) {
	h := &Handler{App: app, Store: s}

	// Setup routes
	h.Init()
}

// Init sets up the routes for the application
func (h *Handler) Init() {
	h.SetupAccount()

	return
}

// SetupAccount sets up the routes for the account resource
func (h *Handler) SetupAccount() *AccountHandler {
	ah := &AccountHandler{h}
	ah.Init()

	return ah
}

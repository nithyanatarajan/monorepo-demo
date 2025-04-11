package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nithyanatarajan/monorepo-demo/services/service-a/api/handlers"
)

// RegisterRoutes registers all API routes
func RegisterRoutes(app *fiber.App, invoiceHandler *handlers.InvoiceHandler) {
	api := app.Group("/api")

	invoices := api.Group("/invoices")
	invoices.Post("/", invoiceHandler.CreateInvoice)
	invoices.Get("/:id", invoiceHandler.GetInvoice)
}

package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nithyanatarajan/monorepo-demo/services/service-b/api/handlers"
)

// RegisterRoutes registers all API routes
func RegisterRoutes(app *fiber.App, webhookHandler *handlers.WebhookHandler) {
	api := app.Group("/api")

	webhooks := api.Group("/webhooks")
	webhooks.Post("/payment", webhookHandler.HandleWebhook)
}

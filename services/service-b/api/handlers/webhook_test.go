package handlers

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/nithyanatarajan/monorepo-demo/pkg/logger"
	"github.com/nithyanatarajan/monorepo-demo/pkg/module-b/payment"
	"github.com/stretchr/testify/assert"
)

func TestWebhookHandler_HandleWebhook(t *testing.T) {
	app := fiber.New()
	handler := NewWebhookHandler(logger.New(), payment.New(logger.New()))

	app.Post("/webhooks", handler.HandleWebhook)

	// Test successful webhook processing
	webhookPayload := `{
		"payment_id": "pay_123",
		"status": "completed",
		"amount": 100.00,
		"currency": "USD"
	}`

	req := httptest.NewRequest("POST", "/webhooks", strings.NewReader(webhookPayload))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// Test invalid request body
	req = httptest.NewRequest("POST", "/webhooks", strings.NewReader("invalid json"))
	req.Header.Set("Content-Type", "application/json")
	resp, err = app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

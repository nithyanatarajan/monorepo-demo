package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/nithyanatarajan/monorepo-demo/pkg/logger"
	"github.com/nithyanatarajan/monorepo-demo/pkg/module-b/payment"
)

// WebhookHandler handles payment webhook requests
type WebhookHandler struct {
	logger  logger.Logger
	payment payment.PaymentGateway
}

// NewWebhookHandler creates a new webhook handler
func NewWebhookHandler(logger logger.Logger, payment payment.PaymentGateway) *WebhookHandler {
	return &WebhookHandler{
		logger:  logger,
		payment: payment,
	}
}

// WebhookRequest represents the payment webhook payload
type WebhookRequest struct {
	PaymentID string  `json:"payment_id"`
	Status    string  `json:"status"`
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
}

// HandleWebhook processes incoming payment webhooks
func (h *WebhookHandler) HandleWebhook(c *fiber.Ctx) error {
	var req WebhookRequest
	if err := c.BodyParser(&req); err != nil {
		h.logger.Error("failed to parse request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Convert request to JSON bytes for processing
	payload, err := json.Marshal(req)
	if err != nil {
		h.logger.Error("failed to marshal request", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process request",
		})
	}

	if err := h.payment.ProcessWebhook(c.Context(), payload); err != nil {
		h.logger.Error("failed to process webhook", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process webhook",
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

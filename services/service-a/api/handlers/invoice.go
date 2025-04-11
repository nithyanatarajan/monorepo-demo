package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/nithyanatarajan/monorepo-demo/pkg/logger"
	"github.com/nithyanatarajan/monorepo-demo/pkg/module-a/billing"
)

// InvoiceHandler handles invoice-related HTTP requests
type InvoiceHandler struct {
	logger  logger.Logger
	service billing.BillingService
}

// NewInvoiceHandler creates a new invoice handler
func NewInvoiceHandler(logger logger.Logger, service billing.BillingService) *InvoiceHandler {
	return &InvoiceHandler{
		logger:  logger,
		service: service,
	}
}

// CreateInvoiceRequest represents the request body for creating an invoice
type CreateInvoiceRequest struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

// CreateInvoice handles the creation of a new invoice
func (h *InvoiceHandler) CreateInvoice(c *fiber.Ctx) error {
	var req CreateInvoiceRequest
	if err := c.BodyParser(&req); err != nil {
		h.logger.Error("failed to parse request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	invoice, err := h.service.CreateInvoice(c.Context(), req.Amount, req.Currency)
	if err != nil {
		h.logger.Error("failed to create invoice", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create invoice",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(invoice)
}

// GetInvoice handles retrieving an invoice by ID
func (h *InvoiceHandler) GetInvoice(c *fiber.Ctx) error {
	invoiceID := c.Params("id")
	if invoiceID == "" {
		h.logger.Error("invoice ID is required", zap.Error(nil))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invoice ID is required",
		})
	}

	invoice, err := h.service.GetInvoice(c.Context(), invoiceID)
	if err != nil {
		h.logger.Error("failed to get invoice", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get invoice",
		})
	}

	return c.JSON(invoice)
}

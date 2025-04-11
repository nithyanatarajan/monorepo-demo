package handlers

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/nithyanatarajan/monorepo-demo/pkg/logger"
	"github.com/nithyanatarajan/monorepo-demo/pkg/module-a/billing"
	"github.com/nithyanatarajan/monorepo-demo/pkg/module-b/payment"
	"github.com/stretchr/testify/assert"
)

func TestInvoiceHandler_CreateInvoice(t *testing.T) {
	app := fiber.New()
	handler := NewInvoiceHandler(logger.New(), billing.New(logger.New(), payment.New(logger.New())))

	app.Post("/invoices", handler.CreateInvoice)

	// Test successful invoice creation
	req := httptest.NewRequest("POST", "/invoices", strings.NewReader(`{"amount": 100.00, "currency": "USD"}`))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

	var invoice billing.Invoice
	body, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &invoice)
	assert.NoError(t, err)
	assert.Equal(t, 100.00, invoice.Amount)
	assert.Equal(t, "USD", invoice.Currency)
}

func TestInvoiceHandler_GetInvoice(t *testing.T) {
	app := fiber.New()
	handler := NewInvoiceHandler(logger.New(), billing.New(logger.New(), payment.New(logger.New())))

	app.Get("/invoices/:id", handler.GetInvoice)

	// Test successful invoice retrieval
	req := httptest.NewRequest("GET", "/invoices/test-id", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var invoice billing.Invoice
	body, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &invoice)
	assert.NoError(t, err)
	assert.Equal(t, "test-id", invoice.ID)
}

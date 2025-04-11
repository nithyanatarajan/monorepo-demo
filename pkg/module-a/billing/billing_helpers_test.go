package billing

import (
	"context"
	"testing"

	"github.com/nithyanatarajan/monorepo-demo/pkg/logger"
	"github.com/nithyanatarajan/monorepo-demo/pkg/module-b/payment"
	"github.com/stretchr/testify/assert"
)

func TestService_CreateInvoice(t *testing.T) {
	service := New(logger.New(), payment.New(logger.New()))
	invoice, err := service.CreateInvoice(context.Background(), 100.00, "USD")
	assert.NoError(t, err)
	assert.NotNil(t, invoice)
	assert.Equal(t, 100.00, invoice.Amount)
	assert.Equal(t, "USD", invoice.Currency)
	assert.Equal(t, "pending", invoice.Status)
}

func TestService_ProcessPayment(t *testing.T) {
	service := New(logger.New(), payment.New(logger.New()))
	err := service.ProcessPayment(context.Background(), "test-id")
	assert.NoError(t, err)
}

func TestService_GetInvoice(t *testing.T) {
	service := New(logger.New(), payment.New(logger.New()))
	invoice, err := service.GetInvoice(context.Background(), "test-id")
	assert.NoError(t, err)
	assert.NotNil(t, invoice)
	assert.Equal(t, "test-id", invoice.ID)
	assert.Equal(t, "paid", invoice.Status)
}

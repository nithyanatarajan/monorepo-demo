package payment

import (
	"context"
	"testing"

	"github.com/nithyanatarajan/monorepo-demo/pkg/logger"
	"github.com/stretchr/testify/assert"
)

func TestSDK_CreatePayment(t *testing.T) {
	sdk := New(logger.New())
	payment, err := sdk.CreatePayment(context.Background(), 100.00, "USD")
	assert.NoError(t, err)
	assert.NotNil(t, payment)
	assert.Equal(t, 100.00, payment.Amount)
	assert.Equal(t, "USD", payment.Currency)
	assert.Equal(t, StatusPending, payment.Status)
}

func TestSDK_GetPayment(t *testing.T) {
	sdk := New(logger.New())
	payment, err := sdk.GetPayment(context.Background(), "test-id")
	assert.NoError(t, err)
	assert.NotNil(t, payment)
	assert.Equal(t, "test-id", payment.ID)
	assert.Equal(t, StatusCompleted, payment.Status)
}

func TestSDK_ProcessWebhook(t *testing.T) {
	sdk := New(logger.New())
	err := sdk.ProcessWebhook(context.Background(), []byte("{}"))
	assert.NoError(t, err)
}

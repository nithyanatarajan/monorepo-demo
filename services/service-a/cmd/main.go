package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/nithyanatarajan/monorepo-demo/pkg/logger"
	"github.com/nithyanatarajan/monorepo-demo/pkg/module-a/billing"
	"github.com/nithyanatarajan/monorepo-demo/pkg/module-b/payment"
	"github.com/nithyanatarajan/monorepo-demo/services/service-a/api/handlers"
	"github.com/nithyanatarajan/monorepo-demo/services/service-a/api/routes"
)

// Build information. Populated at build-time.
var (
	Version   = "dev"
	BuildTime = "unknown"
)

var rootCmd = &cobra.Command{
	Use:   "servicea",
	Short: "Billing Service",
	Long:  "Billing Service for managing invoices and payments",
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the billing service",
	RunE:  serve,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Billing Service %s (Built on %s)\n", Version, BuildTime)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(versionCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func serve(cmd *cobra.Command, args []string) error {
	// Initialize logger
	logger := logger.New()

	// Initialize dependencies
	paymentSDK := payment.New(logger)
	billingService := billing.New(logger, paymentSDK)
	invoiceHandler := handlers.NewInvoiceHandler(logger, billingService)

	// Create Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			logger.Error("request error", zap.Error(err))
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal server error",
			})
		},
	})

	// Register routes
	routes.RegisterRoutes(app, invoiceHandler)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting billing service on :%s", port)
	if err := app.Listen(":" + port); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}

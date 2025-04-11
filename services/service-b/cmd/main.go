package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/nithyanatarajan/monorepo-demo/pkg/logger"
	"github.com/nithyanatarajan/monorepo-demo/pkg/module-b/payment"
	"github.com/nithyanatarajan/monorepo-demo/services/service-b/api/handlers"
	"github.com/nithyanatarajan/monorepo-demo/services/service-b/api/routes"
)

// Build information. Populated at build-time.
var (
	Version   = "dev"
	BuildTime = "unknown"
)

var rootCmd = &cobra.Command{
	Use:   "serviceb",
	Short: "Payment Webhook Service",
	Long:  "Payment Webhook Service for processing payment notifications",
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the payment webhook service",
	RunE:  serve,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Payment Webhook Service %s (Built on %s)\n", Version, BuildTime)
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
	webhookHandler := handlers.NewWebhookHandler(logger, paymentSDK)

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
	routes.RegisterRoutes(app, webhookHandler)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Printf("Starting payment webhook service on :%s", port)
	if err := app.Listen(":" + port); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}

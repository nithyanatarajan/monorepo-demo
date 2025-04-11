# Go Monorepo Demo

This is a demonstration monorepo showcasing a microservices architecture with shared libraries. The project is structured to illustrate:

- Service isolation and boundaries
- Shared library usage
- Dependency management
- Testing organization
- Mock generation

## Project Structure

```
monorepo/
├── services/           # Microservices
│   ├── service-a/     # Billing API service
│   └── service-b/     # Payment Webhook Listener
├── pkg/               # Shared libraries
│   ├── module-a/      # Billing helper library
│   ├── module-b/      # Payment SDK
│   ├── logger/        # Shared logging
│   └── db/            # Database utilities
└── scripts/           # Build and deployment scripts
```

## Services

### Service A: Billing API
- HTTP API for invoice management
- Uses module-a for billing calculations
- Integrates with payment processing

### Service B: Payment Webhook Handler
- Processes payment gateway callbacks
- Uses module-b for payment operations
- Logs transaction results

## API Examples

### Service A: Billing API

#### Create an Invoice
```bash
curl --request POST \
  --url http://localhost:8080/api/invoices \
  --header 'Content-Type: application/json' \
  --data '{
    "customer_id": "cust_123",
    "amount": 100.00,
    "currency": "USD",
    "description": "Monthly subscription"
  }'
```

#### Get an Invoice
```bash
curl --request GET \
  --url http://localhost:8080/api/invoices/inv_123
```

### Service B: Payment Webhook Handler

#### Send a Payment Webhook Notification
```bash
curl --request POST \
  --url http://localhost:8081/api/webhooks/payment \
  --header 'Content-Type: application/json' \
  --data '{
    "payment_id": "payment_123",
    "status": "completed",
    "amount": 100.00,
    "currency": "USD"
  }'
```

## Shared Libraries

### Module A: Billing Helpers
- Billing calculation utilities
- Depends on module-b for payments

### Module B: Payment SDK
- Payment gateway integration
- Used by both services

### Logger
- Shared logging utilities
- Used across all services and modules

### DB
- Database access utilities
- Common database operations

## Development

This is a demonstration project and not intended for production use. The code contains placeholder implementations to illustrate structure and relationships. 

## Available Tasks

The project uses [Task](https://taskfile.dev/) for managing common development tasks. Here are the available commands:

```bash
# Install development dependencies (including linter)
task deps

# Run all tests
task test

# Run linter
task lint

# Build all services
task build

# Clean build artifacts
task clean
```

Each service will be built into the `bin/` directory. The services can be run directly from there:

```bash
# Run Service A (Billing API)
./bin/service-a serve

# Run Service B (Payment Webhook Handler)
./bin/service-b serve
```
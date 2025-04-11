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
- Runs on port 8080

### Service B: Payment Webhook Handler
- Processes payment gateway callbacks
- Uses module-b for payment operations
- Logs transaction results
- Runs on port 8081

## API Examples

### Service A: Billing API

#### Create an Invoice
```sh
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
```sh
curl --request GET \
  --url http://localhost:8080/api/invoices/inv_123
```

### Service B: Payment Webhook Handler

#### Send a Payment Webhook Notification
```sh
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
- Used by Service A

### Module B: Payment SDK
- Payment gateway integration
- Used by both services
- Provides payment processing functionality

### Logger
- Shared logging utilities
- Used across all services and modules
- Provides consistent logging format

### DB
- Database access utilities
- Common database operations
- Connection pooling and management

## Development

This is a demonstration project and not intended for production use. The code contains placeholder implementations to illustrate structure and relationships. 

## Available Tasks

The project uses [Task](https://taskfile.dev/) for managing common development tasks. The tasks are convention-based and will automatically discover any services added to the project.

### Development Setup

```sh
# Install development dependencies (including linter)
task deps
```

### Common Tasks

```sh
# Run all tests
task test

# Run tests with coverage report
task test-coverage

# Run linter
task lint

# Build all services
task build

# Build a specific service
task build-service SERVICE=service-a

# Run a specific service
task run SERVICE=service-a

# List available services
task list

# Clean build artifacts and coverage reports
task clean
```

### Service Conventions

The project follows these conventions for services:

1. All services must be placed in the `services/` directory
2. Each service must have a `cmd/` directory containing its main package
3. Services are built into the `bin/` directory with their directory name as the binary name

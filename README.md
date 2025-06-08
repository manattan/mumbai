# mumbai

## Project Structure

```
.
├── cmd/
│   ├── echo-server/           # HTTP server entry point
│   └── grpc-server/           # gRPC server entry point
├── configs/
│   └── dev.env               # Development environment variables
├── internal/
│   ├── config/               # Configuration management
│   ├── domain/               # Domain layer (entities, interfaces)
│   ├── gateway/              # Infrastructure layer
│   ├── middleware/           # HTTP middleware
│   ├── response/             # Response utilities
│   └── usecase/              # Use case implementations
├── handler/                  # HTTP handlers
├── usecase/                  # Use case interfaces and implementations
├── migrations/               # Database migration files
├── pkg/                      # Shared packages
└── test/                     # Test files
```

## Commands

### Development

```bash
# Run the HTTP server
go run cmd/echo-server/main.go

# Run tests
go test ./...
```

### Docker

```bash
# Build and start services
docker-compose up -d

# Stop services
docker-compose down
```
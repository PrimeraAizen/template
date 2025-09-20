# Go Clean Architecture Template

A production-ready Go web application template built with clean architecture principles, featuring Gin web framework, PostgreSQL database, and comprehensive project structure.

## 🏗️ Architecture

This template follows **Clean Architecture** principles with clear separation of concerns:

```
├── cmd/                    # Application entry points
│   └── web/               # Web server main
├── config/                # Configuration management
├── internal/              # Private application code
│   ├── app/              # Application initialization
│   ├── delivery/         # Delivery layer (HTTP handlers)
│   │   ├── dto/          # Data Transfer Objects
│   │   └── rest/         # REST API handlers
│   ├── domain/           # Domain entities and business rules
│   ├── repository/       # Data access layer
│   ├── server/           # HTTP server configuration
│   └── service/          # Business logic layer
├── migrations/           # Database migrations
├── pkg/                  # Public packages
│   └── adapter/          # External service adapters
└── config/               # Configuration files
```

## 🚀 Features

- **Clean Architecture** with proper layer separation
- **Gin Web Framework** for high-performance HTTP routing
- **PostgreSQL** with connection pooling
- **Database Migrations** using Goose
- **Structured Logging** with JSON output
- **Configuration Management** with Viper
- **Docker Support** with multi-stage builds
- **Health Checks** for monitoring
- **Graceful Shutdown** handling

## 📋 Prerequisites

- Go 1.23.0 or higher
- PostgreSQL 12 or higher
- Docker (optional, for containerized deployment)

## 🛠️ Installation

1. **Clone the repository:**
   ```bash
   git clone <your-repo-url>
   cd my-template
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Set up configuration:**
   ```bash
   cp config/config.example.yaml config/config.yaml
   # Edit config/config.yaml with your database settings
   ```

4. **Set up the database:**
   ```bash
   # Create PostgreSQL database
   createdb your_database_name
   
   # Run migrations
   make migrate-up DB_URL="postgres://username:password@localhost:5432/your_database_name"
   ```

## 🏃‍♂️ Running the Application

### Development Mode

```bash
# Run with hot reload (requires air or similar tool)
make run

# Or run directly
go run cmd/web/main.go
```

### Production Mode

```bash
# Build the application
make build

# Run the binary
./bin/myapp
```

### Docker

```bash
# Build Docker image
docker build -t my-template .

# Run with Docker Compose (if available)
docker-compose up
```

## 🔧 Configuration

The application uses YAML configuration files with environment variable overrides:

### Configuration Structure

```yaml
http:
  host: localhost
  port: "8080"

database:
  host: localhost
  port: "5432"
  database: postgres
  username: postgres
  password: change-me
  ssl_mode: disable
  max_conns: 10
  min_conns: 1
```

### Environment Variables

You can override any configuration value using environment variables with the `APP_` prefix:

```bash
export APP_HTTP_HOST=0.0.0.0
export APP_HTTP_PORT=8080
export APP_DATABASE_HOST=localhost
export APP_DATABASE_PASSWORD=your_password
```

## 📊 API Endpoints

### Health Checks

- `GET /v1/healthz` - Basic health check
- `GET /v1/readyz` - Readiness check (includes database connectivity)

### Example Endpoints

- `GET /v1/example/` - Example endpoint demonstrating the architecture

## 🗄️ Database Migrations

This template uses [Goose](https://github.com/pressly/goose) for database migrations.

### Available Commands

```bash
# Create a new migration
make migrate-new name=create_users_table

# Apply migrations
make migrate-up DB_URL="postgres://user:pass@localhost/db"

# Rollback migrations
make migrate-down DB_URL="postgres://user:pass@localhost/db"

# Check migration status
make migrate-status DB_URL="postgres://user:pass@localhost/db"
```

## 🏗️ Project Structure Details

### Domain Layer (`internal/domain/`)
- Contains business entities and rules
- Independent of external frameworks
- Defines interfaces for repositories and services

### Repository Layer (`internal/repository/`)
- Implements data access interfaces
- Handles database operations
- Uses Squirrel for query building

### Service Layer (`internal/service/`)
- Contains business logic
- Orchestrates repository calls
- Implements domain interfaces

### Delivery Layer (`internal/delivery/`)
- HTTP handlers and middleware
- Request/response transformation
- Input validation

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./internal/service/...
```

## 📦 Building

```bash
# Build for current platform
make build

# Build for specific platform
GOOS=linux GOARCH=amd64 go build -o bin/myapp-linux cmd/web/main.go

# Clean build artifacts
make clean
```

## 🐳 Docker

The project includes a multi-stage Dockerfile for optimized production builds:

```dockerfile
FROM golang:1.23 AS build
# ... build stage

FROM gcr.io/distroless/base-debian12
# ... minimal runtime stage
```

### Docker Commands

```bash
# Build image
docker build -t my-template .

# Run container
docker run -p 8080:8080 my-template

# Run with environment variables
docker run -p 8080:8080 \
  -e APP_DATABASE_HOST=host.docker.internal \
  -e APP_DATABASE_PASSWORD=your_password \
  my-template
```

## 📈 Monitoring & Observability

### Health Checks

The application provides health check endpoints for monitoring:

- `/v1/healthz` - Basic application health
- `/v1/readyz` - Application readiness (includes database connectivity)

### Logging

Structured JSON logging is configured by default:

```json
{
  "time": "2024-01-15T10:30:00Z",
  "level": "INFO",
  "msg": "Web server started!",
  "source": "app.go:35"
}
```

## 🔒 Security Considerations

- Database credentials should be managed via environment variables
- Use SSL/TLS in production environments
- Implement proper authentication and authorization
- Add rate limiting and request validation
- Use security headers and CORS configuration

## 🚀 Deployment

### Environment Setup

1. **Production Database:**
   - Set up PostgreSQL with proper security
   - Configure connection pooling
   - Enable SSL connections

2. **Configuration:**
   - Use environment variables for sensitive data
   - Configure proper logging levels
   - Set up monitoring and alerting

3. **Container Deployment:**
   - Use the provided Dockerfile
   - Configure health checks
   - Set up proper resource limits

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🆘 Support

If you encounter any issues or have questions:

1. Check the [Issues](https://github.com/your-username/my-template/issues) page
2. Create a new issue with detailed information
3. Provide logs and configuration details

## 🔄 Version History

- **v1.0.0** - Initial template with basic clean architecture
- **v1.1.0** - Added health checks and improved configuration
- **v1.2.0** - Enhanced Docker support and logging

---

**Happy Coding! 🎉**

APP_NAME=myapp

.PHONY: run build clean

# Run the application
run:
	go run cmd/web/main.go

# Build the application binary
build:
	go build -o bin/$(APP_NAME) cmd/web/main.go

# Clean build artifacts
clean:
	rm -rf bin

# Create a new migration file
migrate-new:
	goose -dir migrations create $(name) sql

# Apply migrations
migrate-up:
	goose -dir migrations -table goose_db_version postgres "$(DB_URL)" up

# Cancel migrations
migrate-down:
	goose -dir migrations -table goose_db_version postgres "$(DB_URL)" down

migrate-status:
	goose -dir migrations -table goose_db_version postgres "$(DB_URL)" status


FRONTEND_DIR=../frontend

.PHONY: frontend backend up down downdb

frontend:
	@cd $(FRONTEND_DIR) && sh -c '$(CMD)'

up:
	@echo "Starting frontend service..."
	@$(MAKE) --no-print-directory frontend CMD="docker compose up -d"

up-build:
	@echo "Building and starting frontend service..."
	@$(MAKE) --no-print-directory frontend CMD="docker compose up -d --build"

down:
	@echo "Killing frontend service..."
	@$(MAKE) --no-print-directory frontend CMD="docker compose down"

downdb:
	@echo "Killing frontend service & wiping database..."
	@$(MAKE) --no-print-directory frontend CMD="docker compose down --volumes"

backend:
	@echo "Starting backend service..."
	@go run ./cmd/api/main.go

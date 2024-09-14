BACKEND_DIR=../backend

.PHONY: backend frontend up down downdb

backend:
	@cd $(BACKEND_DIR) && sh -c  '$(CMD)'

up:
	@echo "Starting backend service..."
	@$(MAKE) --no-print-directory backend CMD="docker compose up -d"

up-build:
	@echo "Building and starting backend service..."
	@$(MAKE) --no-print-directory backend CMD="docker compose up -d --build"

down:
	@echo "Killing backend service..."
	@$(MAKE) --no-print-directory backend CMD="docker compose down"

down-db:
	@echo "Killing backend service & wiping database..."
	@$(MAKE) --no-print-directory backend CMD="docker compose down --volumes"

frontend:
	@echo "Starting frontend service..."
	@ng serve

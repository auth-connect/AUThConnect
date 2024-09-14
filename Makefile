.PHONY: all

up:
	@echo "Starting all services..."
	@docker compose up -d

up-build:
	@echo "Building and starting all services..."
	@docker compose up -d --build

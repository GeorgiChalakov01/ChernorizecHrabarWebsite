.PHONY: clear build down up restart

clear:
	@clear

build:
	@echo "Rebuilding backend..."
	@cd app && make generate build
	./build.sh
	@echo "Updating backend container..."
	@docker compose up -d --no-deps backend

down:
	@docker compose down

up:
	@docker compose up -d

restart: down up

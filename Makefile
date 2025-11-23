ifneq (,$(wildcard ./.env))
	include .env
	export
endif


MIGRATIONS_PATH=./database/migrations

DATABASE_URL=postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DATABASE)?sslmode=disable

migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "Usage: make migrate-create name=migration_name"; \
	else \
		migrate create -ext sql -dir $(MIGRATIONS_PATH) "$(name)"; \
	fi

migrate-version:
	@migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" version

migrate-force:
	@if [ -z "$(version)" ]; then \
		echo "Usage: make migrate-force version=<migration_version>"; \
		echo "Example: make migrate-force version=20251122105423"; \
		echo "Note: <migration_version> is the timestamp prefix in migration file names."; \
	else \
		migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" force $(version); \
	fi

migrate-goto:
	@if [ -z "$(version)" ]; then \
		echo "Usage: make migrate-goto version=<migration_version>"; \
		echo "Example: make migrate-goto version=20251122105423"; \
		echo "Note: <migration_version> is the timestamp prefix in migration file names."; \
	else \
		migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" goto $(version); \
	fi

migrate-up:
	@if [ -z "$(step)" ]; then \
		migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" up; \
	else \
		migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" up $(step); \
	fi

migrate-down:
	@if [ -z "$(step)" ]; then \
		migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" down; \
	else \
		migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" down $(step); \
	fi

help:
	@echo ""
	@echo "Available commands:"
	@echo ""
	@echo "  migrate-version    Show current migration version"
	@echo "      Usage: make migrate-version"
	@echo ""
	@echo "  migrate-create     Create new migration file (requires: name=)"
	@echo "      Usage: make migrate-create name=migration_name"
	@echo ""
	@echo "  migrate-up         Run migration up (optional: step=)"
	@echo "      Usage: make migrate-up"
	@echo "      Usage: make migrate-up step=1"
	@echo ""
	@echo "  migrate-down       Run migration down (optional: step=)"
	@echo "      Usage: make migrate-down"
	@echo "      Usage: make migrate-down step=1"
	@echo ""
	@echo "  migrate-force      Force set migration version (requires: version=)"
	@echo "      NOTE: 'version' must match the timestamp prefix of migration files."
	@echo "      Example filename: 20251122105423_create_table.up.sql"
	@echo "      Usage: make migrate-force version=20251122105423"
	@echo ""
	@echo "  migrate-goto       Migrate database to a specific version (runs up or down)"
	@echo "      NOTE: 'version' must match the timestamp prefix of migration files."
	@echo "      Usage: make migrate-goto version=20251122105423"
	@echo ""

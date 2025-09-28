# Load ENV
ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

# Default environment
GO_ENV ?= local

# Set sslmode based on GO_ENV
ifeq ($(GO_ENV),local)
	SSL_MODE=disable
else
	SSL_MODE=require
endif

DB_URL = postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE)
MIGRATIONS_PATH = migrations

# -----------------------
# Migration commands
# -----------------------

migrate-up:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

migrate-down:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down 1

migrate-force:
	@echo "Enter version to force:"; \
	read version; \
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" force $$version

migrate-reset:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

create-migration:
	@echo "Enter migration name:"; \
	read name; \
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $$name

.PHONY: migrate-up migrate-down migrate-force migrate-reset create-migration


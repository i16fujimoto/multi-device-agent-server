# ==========================
# Docker
# ==========================
.PHONY: build
build:
	docker-compose build $(ARGS)
	docker-compose up -d

.PHONY: up
up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down

# ==========================
# go mod
# ==========================
.PHONY: go-mod-tidy
go-mod-tidy:
	go1.22.4 mod tidy

# ==========================
# Lint
# ==========================

.PHONY: lint
lint: ## Run lint
	golangci-lint run ./...
.PHONY: lint/fix
lint/fix: ## Run lint and fix
	golangci-lint run --fix ./...

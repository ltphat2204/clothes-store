include .env

.PHONY: server
server:
	@echo "Starting server..."
	@go run ./cmd/app/

.PHONY: client
client:
	@echo "Starting client..."
	@cd web && npm run dev

.PHONY postgres:
postgres:
	@echo "Starting postgres..."
	@docker run --name postgres-db -p 5432:5432 -e POSTGRES_USER="${DATABASE_USER}" -e POSTGRES_PASSWORD="${DATABASE_PASSWORD}" -d postgres:16-alpine

.PHONY database:
database:
	@echo "Starting create database ${DATABASE_NAME}..."
	@docker exec -it postgres-db createdb --username="${DATABASE_USER}" --owner="${DATABASE_USER}" "${DATABASE_NAME}"

.PHONY setup:
setup: postgres database
.PHONY: server
server:
	@echo "Starting server..."
	@go run ./cmd/app/

.PHONY: client
client:
	@echo "Starting client..."
	@cd web && npm run dev
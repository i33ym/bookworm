build:
	@echo "Building the application..."
	@go build -o ./bin/bookworm-api ./cmd/api

run: build
	@./bin/bookworm-api
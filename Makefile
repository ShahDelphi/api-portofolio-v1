.PHONY: help run build seed clean

help:
	@echo "Available commands:"
	@echo "  make run      - Run the application"
	@echo "  make build    - Build the application"
	@echo "  make seed     - Run database seeder"
	@echo "  make clean    - Clean build files"

run:
	go run cmd/api/main.go

build:
	go build -o bin/portfolio-api.exe cmd/api/main.go

seed:
	go run cmd/api/main.go --seed

clean:
	rm -rf bin/
	go clean

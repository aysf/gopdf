test:
	@echo "testing code..."
	go test -v ./...

api:
	@echo "starting web api..."
	go run cmd/api/*.go
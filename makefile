test:
	@echo "testing code..."
	go test -v ./...

testc:
	@echo "testing with coverage..."
	go test -v -race -coverprofile=coverage.out -covermode=atomic -failfast ./...

api:
	@echo "starting web api..."
	go run cmd/api/*.go
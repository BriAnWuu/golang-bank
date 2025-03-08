build:
	@go build -o bin/golang-bank

run: build
	@./bin/golang-bank

test:
	@go test -v ./...
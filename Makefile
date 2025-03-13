build:
	@go build -o bin/golang-bank

run: build
	@./bin/golang-bank

run-seed: build
	@./bin/golang-bank --seed

test:
	@go test -v ./...
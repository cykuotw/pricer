build:
	@go mod tidy
	@go build -o bin/pricer main.go

test: 
	@go test -v ./...

run: build
	@./bin/pricer
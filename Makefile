build:
	@go build -o bin/wapi

run: build
	@./bin/wapi

test:
	@go test -v ./...
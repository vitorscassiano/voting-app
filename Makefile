.PHONY: run tests

run:
	@go run main.go

test:
	@go test ./...
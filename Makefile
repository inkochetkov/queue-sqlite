.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: mod
mod:
	go mod tidy

.PHONY: linter-install
linter-install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: linter
linter: 
	golangci-lint run ./...
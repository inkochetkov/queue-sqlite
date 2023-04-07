.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: mod
mod:
	go mod tidy

setup:
	go mod download && \
	go mod tidy

lint:
	hadolint Dockerfile
	go vet ./...

test:
	go clean -testcache ./... && \
	go test -v -race ./...

all: setup lint test
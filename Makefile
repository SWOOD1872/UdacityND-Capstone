setup:
	go mod download && \
	go mod tidy

lint:
	hadolint Dockerfile
    golangci-lint run ./...

vet:
	go vet ./...

test:
	go clean -testcache ./... && \
	go test -v -race ./...

deploy:
	kubectl apply -f kubernetes.yml
	kubectl get all

all: setup lint vet test deploy
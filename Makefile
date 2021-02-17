test:
	go clean -testcache ./... && \
	go test -v -race ./...

all: test
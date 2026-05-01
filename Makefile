build:
	@go build -o bin/dev cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/dev

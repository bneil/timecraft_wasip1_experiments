GO ?= gotip

run:
	timecraft run app.wasm

build:
	GOOS=wasip1 GOARCH=wasm $(GO) build -o app.wasm main.go

lint:
	golangci-lint run ./...
	buf lint

fmt:
	$(GO) fmt ./...

.PHONY: run test fmt vet tidy

run:
	go run ./cmd/gateway

test:
	go test ./...

fmt:
	gofmt -w $$(find . -name '*.go' -not -path './vendor/*')

vet:
	go vet ./...

tidy:
	go mod tidy

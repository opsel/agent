default: run

run:
	@go run agent.go

build:
	@go build -o opsel -ldflags "-s -w" -trimpath agent.go

test:
	@go test

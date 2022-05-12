runServer:
	wire ./cmd/server/wire.go
	go run ./cmd/server

runClient:
	go run ./cmd/client

fetch:
	go mod download
	go mod verify

.DEFAULT_GOAL := runServer
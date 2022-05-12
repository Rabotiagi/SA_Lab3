runServer:
	wire ./cmd/server/wire.go
	go run ./cmd/server

runClient:
	go run ./cmd/client

.DEFAULT_GOAL := runServer
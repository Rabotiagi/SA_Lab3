runServer:
	go run ./cmd/server

runClient:
	go run ./cmd/client

.DEFAULT_GOAL := runServer
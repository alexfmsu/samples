generate = bash generate.sh

install:
	@go mod init server_streaming 2>/dev/null || true
	@go mod tidy

run:
ifeq (server, $(filter server, $(MAKECMDGOALS)))
	@${generate}
	@go run server/server.go
else ifeq (client, $(filter client, $(MAKECMDGOALS)))
	@${generate}
	@go run client/client.go
endif

.PHONY: server client

server:
	@:

client:
	@:

clean:
	@rm -rf go.mod go.sum 2>/dev/null proto/lotspb || true

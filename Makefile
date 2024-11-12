BUILD_DIR=bin/build/
S_BINARY=server
C_BINARY=client
S_BIN_DIR=$(BUILD_DIR)$(S_BINARY)
C_BIN_DIR=$(BUILD_DIR)$(C_BINARY)

build:
	go build -o $(S_BIN_DIR)

build-client:
	cd ./pkg/client && go build -o ../../$(C_BIN_DIR)

build-all: build build-client

run:
	go run main.go

run-build:
	./$(S_BIN_DIR)

run-client:
	cd ./pkg/client && go run client.go

run-client-build:
	./$(C_BIN_DIR)

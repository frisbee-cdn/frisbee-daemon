.PHONY: help build compile-proto test test_integration

SOURCE_DIR = $$(pwd)

PROTO_LOCATION = $(SOURCE_DIR)/pkg/rpc/proto

help:
	@echo "make help               Show this help"
	@echo "make build              Build the binaries in build directory"
	@echo "make test               Run the unit tests"
	@echo "make compile-proto      Compiles protocol-buffers resources"
	@echo "make test-integration   Run the all the tests, including integration tests"

build: test
	GOOS=linux GOARCG=amd64 go build -o ./build/linux_amd64/opype .
	GOOS=darwin GOARCG=amd64 go build -o ./build/darwin_amd64/opype . 
	GOOS=windows GOARCG=amd64 go build -o ./build/windows_amd64/opype.exe .

compile-proto:
	protoc -I=$(PROTO_LOCATION) --go_out=plugins=grpc:$(PROTO_LOCATION) \
			$(PROTO_LOCATION)/frisbee.proto

test_integration: test
	go test -tags=integration


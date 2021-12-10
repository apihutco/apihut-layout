GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
API_PROTO_FILES=$(shell find api -name *.proto)
PROJECT_NAME = $(shell basename $(shell pwd))

.PHONY: init
# init env
init:
	go get -u github.com/go-kratos/kratos/cmd/kratos/v2
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2
	go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2


.PHONY: api
# generate api proto
api:
	protoc --proto_path=. \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
 	       --go-http_out=paths=source_relative:. \
 	       --go-grpc_out=paths=source_relative:. \
 	       --validate_out=paths=source_relative,lang=go:. \
 	       --go-errors_out=paths=source_relative:. \
	       $(API_PROTO_FILES)


.PHONY: config
# generate internal proto
config:
	protoc --proto_path=. \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
	       $(INTERNAL_PROTO_FILES)


.PHONY: ent
# generate ent
ent:
	go generate ./internal/data/ent

.PHONY: wire
# generate ent
wire:
	wire ./cmd/$(PROJECT_NAME)

.PHONY: server
# generate server in (internal/service)
server:
	kratos proto server $(API_PROTO_FILES) -t internal/service


.PHONY: generate
# equal go generate ./...
generate:
	go generate ./...


.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: run
# build and run
run:
	make build && cd bin && ./$(PROJECT_NAME).exe -conf ../configs

.PHONY: crun
# remove /bin and build and run
crun:
	rm -rf bin && make run


.PHONY: all
# generate all
all:
	make api;
	make config;
	make generate;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

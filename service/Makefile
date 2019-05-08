NAME       := grpc-bridge
REVISION   := $(shell git rev-parse --short HEAD)

SRCS    := $(shell find . -type f -name '*.go')
LDFLAGS := -ldflags="-s -w -X \"main.Revision=$(REVISION)\" -extldflags \"-static\""
SQ_ENV ?= test

bin/$(NAME): $(SRCS)
	@go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o bin/$(NAME)

.PHONY: install
install:
	go install $(LDFLAGS)

.PHONY: mod-download
mod-download:
	go mod download

.PHONY: go-test
go-test:
	@docker-compose run test go test ./...

.PHONY: go-vet
go-vet:
	@docker-compose run test go vet ./...

.PHONY: docker-builder-build
docker-builder-build:
	docker build --target builder -t $(NAME):builder .

.PHONY: docker-build
docker-build:
	docker build --cache-from=$(NAME):builder -t $(NAME):1.0.0 .

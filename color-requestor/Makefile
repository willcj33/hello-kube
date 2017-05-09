SHELL := /bin/bash

#BASIC COMMANDS
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_BUILD_RACE=$(GO_CMD) build -race
GO_TEST=$(GO_CMD) test
GO_TEST_VERBOSE=$(GO_CMD) test -v
GO_INSTALL=$(GO_CMD) install -v
GO_CLEAN=$(GO_CMD) clean
GO_DEPS=$(GO_CMD) get -d -v
GO_DEPS_UPDATE=$(GO_CMD) get -d -v -u
GO_VET=$(GO_CMD) vet
GO_FMT=$(GO_CMD) fmt
GO_LINT=golint
SERVICE_NAME=color-requestor


#DIR STUFF
TOP_PACKAGE_DIR := github.com/willcj33/hello-kube
TEST_DIR := test
VENDOR_PATH := vendor/github.com/willcj33

#PROTO
PROTOC_CMD=protoc
PROTOC_GO_OUT=--go_out
PROTOC_PLUGINS=plugins
PROTOC_GRPC=grpc
VENDOR_PROTO_PATH=$(VENDOR_PATH)/proto-test
PROTO_LIST=charlotte/common/traits/person/v1/*.proto charlotte/common/source/v1/*.proto charlotte/common/traits/profile/v1/*.proto charlotte/common/profile/v1/*.proto charlotte/common/person/v1/*.proto
LOCAL_PROTO_PATH=../../../../proto/
GLOBAL_PROTO_REPO=github.com/willcj33/proto-test

#GLIDE
GLIDE_CMD=glide
GLIDE_CLEAN=$(GLIDE_CMD) cc
GLIDE_UPDATE=$(GLIDE_CMD) update
GLIDE_INSTALL=$(GLIDE_CMD) install
GLIDE_CLEAN_AND_INSTALL=$(GLIDE_CLEAN) && $(GLIDE_UPDATE) && $(GLIDE_INSTALL)

#DOCKER VARS
PROJECT=color-requestor
VERSION=0.1


.PHONY: all build build-race test test-verbose deps update-deps install clean fmt vet lint

all: glide-pre-fetch protoc glide build test-verbose

build: vet
	@for p in $(SERVICE_NAME); do \
		echo "==> Build $$p ..."; \
		$(GO_BUILD) -o ./bin/$(PROJECT) || exit 1; \
	done

build-linux:
	@for p in $(SERVICE_NAME); do \
		echo "==> Build $$p ..."; \
		GOOS=linux CGO_ENABLED=0 go build -ldflags "-s" -a -installsuffix cgo -o ./bin/$(PROJECT) || exit 1; \
	done

build-race: vet
	@for p in $(SERVICE_NAME); do \
		echo "==> Build $$p ..."; \
		$(GO_BUILD_RACE) $(TOP_PACKAGE_DIR)/$$p || exit 1; \
	done

test: deps
	@for p in $(SERVICE_NAME); do \
		echo "==> Unit Testing $$p ..."; \
		$(GO_TEST) $(shell go list ./... | grep -v /vendor/) || exit 1; \
	done

test-verbose: deps
	@for p in $(SERVICE_NAME); do \
		echo "==> Unit Testing $$p ..."; \
		$(GO_TEST_VERBOSE) $(shell go list ./... | grep -v /vendor/) || exit 1; \
	done

deps:
	@for p in $(SERVICE_NAME); do \
		echo "==> Install dependencies for $$p ..."; \
		$(GO_DEPS) $(TOP_PACKAGE_DIR)/$$p || exit 1; \
	done

update-deps:
	@for p in $(SERVICE_NAME); do \
		echo "==> Update dependencies for $$p ..."; \
		$(GO_DEPS_UPDATE) $(TOP_PACKAGE_DIR)/$$p || exit 1; \
	done

install:
	@for p in $(SERVICE_NAME); do \
		echo "==> Install $$p ..."; \
		$(GO_INSTALL) $(TOP_PACKAGE_DIR)/$$p || exit 1; \
	done

clean:
	@for p in $(SERVICE_NAME); do \
		echo "==> Clean $$p ..."; \
		$(GO_CLEAN) $(TOP_PACKAGE_DIR)/$$p; \
	done

fmt:
	@for p in $(SERVICE_NAME); do \
		echo "==> Formatting $$p ..."; \
		$(GO_FMT) $(TOP_PACKAGE_DIR)/$$p || exit 1; \
	done
vet:
	@for p in $(SERVICE_NAME); do \
		echo "==> Vet $$p ..."; \
		$(GO_VET) $(TOP_PACKAGE_DIR)/$$p; \
	done

lint:
	@for p in $(SERVICE_NAME); do \
		echo "==> Lint $$p ..."; \
		$(GO_LINT) src/$(TOP_PACKAGE_DIR)/$$p; \
	done

glide:
	@for p in $(SERVICE_NAME); do \
		echo "==> Glide $$p ..."; \
		$(GLIDE_CLEAN_AND_INSTALL); \
	done
	cp -rf ./proto/* ./vendor


glide-pre-fetch:
	@for p in $(SERVICE_NAME); do \
		echo "==> Get Protos $$p ..."; \
		$(GLIDE_CMD) --yaml glide-init.yaml update && $(GLIDE_CMD) --yaml glide-init.yaml install; \
	done

protoc: glide-pre-fetch
	rm -rf ./proto
	mkdir ./proto
	@for p in $(PROTO_LIST); do \
		echo "==> Protoc $$p ..."; \
		pushd ./$(VENDOR_PROTO_PATH);\
	    $(PROTOC_CMD) $(PROTOC_GO_OUT)=$(PROTOC_PLUGINS)=$(PROTOC_GRPC):$(LOCAL_PROTO_PATH) $$p;\
		popd;\
	done


build-docker-local: glide build-linux
	docker build -f Dockerfile.local -t color-requestor:local-dev .
	echo "\nDgs color-requestor service built. It is now available to run in minikube.\n"

.DEFAULT_GOAL := help
.PHONY : build

GOPATH := $(shell go env GOPATH)
DESTDIR?=
PREFIX?=/usr/local
INSTALL_DIR?=$(GOPATH)/bin
MAN_DIR?=$(PREFIX)/share/man

APP := devto
VERSION := $(shell git describe --always)

RFC_3339 := "+%Y-%m-%dT%H:%M:%SZ"
DATE := $(shell date -u $(RFC_3339))
COMMIT := $(shell git rev-list -1 HEAD)

OPTS?=GO111MODULE=on

DEVTO_REPO := github.com/Gealber/devto-cli
BUILDINFO_PATH := $(DEVTO_REPO)/buildinfo

BUILDINFO_VERSION := -X $(BUILDINFO_PATH).version=$(VERSION)
BUILDINFO_DATE := -X $(BUILDINFO_PATH).date=$(DATE)
BUILDINFO_COMMIT := -X $(BUILDINFO_PATH).commit=$(COMMIT)


BUILDINFO?=$(BUILDINFO_VERSION) $(BUILDINFO_DATE) $(BUILDINFO_COMMIT)

BUILD_OPTS?="-ldflags=$(BUILDINFO)"
BUILD_OPTS_DEPLOY?="-ldflags=$(BUILDINFO) -w -s"


build: ## Build binary
	go build ${BUILD_OPTS} -o ${APP}

manpage: ## Create manpage. Need sudo permissions
	mkdir -p $(MAN_DIR)/man1
	install -m 644 share/man/devto.1 $(MAN_DIR)/man1/devto.1

install: ## Install binary and allow manpage
	go build $(BUILD_OPTS) -o $(INSTALL_DIR)/$(APP)

clean: ## Cleaning binary
	rm -f ${APP}

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

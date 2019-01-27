# variable definitions
NAME := lab
DESC := Command line interface to Gitlab
VERSION := $(shell git describe --tags --always --dirty)
SHA := $(shell git rev-parse HEAD)
GOVERSION := $(shell go version)
BUILDTIME := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
BUILDDATE := $(shell date -u +"%B %d, %Y")
BUILDER ?= $(shell echo "`git config user.name` <`git config user.email`>")
TARGET := x86_64-linux
GOOS := linux
PROJECT_URL := "https://github.com/joshbohde/$(NAME)"
GOPACKAGE := "github.com/joshbohde/$(NAME)"
LDFLAGS := -X 'main.version=$(VERSION)' \
           -X 'main.buildTime=$(BUILDTIME)' \
           -X 'main.builder=$(BUILDER)' \
           -X 'main.goversion=$(GOVERSION)' \
           -X 'main.target=$(TARGET)'

.PHONY: dependencies test dist all clean

all: dependencies test dist

test:
	go test -v ./...

dependencies:
	go mod download

CMD_SOURCES := $(shell find cmd -name main.go)
TARGETS := $(patsubst cmd/%/main.go,build/%,$(CMD_SOURCES))

build:
	mkdir -p build

build/%: cmd/%/main.go | build
	GOOS=${GOOS} go build -ldflags "$(LDFLAGS)" -o $@ $<

dist: $(TARGETS)

install: $(TARGETS)
	cp ${TARGETS} /go/bin

clean:
	rm -rf ./build

# variable definitions
NAME := lab
DESC := Command line interface to Gitlab
VERSION := $(shell git describe --tags --always --dirty)
SHA := $(shell git rev-parse HEAD)
GOVERSION := $(shell go version)
BUILDTIME := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
BUILDDATE := $(shell date -u +"%B %d, %Y")
BUILDER ?= $(shell echo "`git config user.name` <`git config user.email`>")
PROJECT_URL := https://github.com/joshbohde/$(NAME)
GOPACKAGE := github.com/joshbohde/$(NAME)
LDFLAGS := -X 'main.version=$(VERSION)' \
           -X 'main.buildTime=$(BUILDTIME)' \
           -X 'main.builder=$(BUILDER)' \
           -X 'main.goversion=$(GOVERSION)'

.PHONY: dependencies test dist all clean

all: dependencies test dist

test:
	go test -v ./...

dependencies:
	go mod download

CMD_SOURCES := $(shell find cmd -maxdepth 1 -mindepth 1 -type d)
TARGETS := $(patsubst cmd/%,target/%,$(CMD_SOURCES))

# For testing out build flow locally
%: cmd/% **/*.go | build
	go build -ldflags "$(LDFLAGS)" -o $@ $(GOPACKAGE)/$<

target/%: cmd/% *.go **/*.go
	mkdir -p $@
	CGO_ENABLED=0 gox -os="linux darwin windows" -arch="amd64" -output="target/$*/$*_{{.OS}}_{{.Arch}}" -ldflags "$(LDFLAGS)" -verbose $(GOPACKAGE)/$<

dist: $(TARGETS)

clean:
	rm -rf ./target

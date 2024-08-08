# Makefile
BINARY_NAME=ezeth
VERSION_PACKAGE=github.com/hakai-here/scylla/constants
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')

all: build

build:
	go build -ldflags "-X $(VERSION_PACKAGE).Version=$(VERSION)"  -o $(BINARY_NAME) cmd/main.go
clean:
	if [ -f $(BINARY_NAME) ] ; then rm $(BINARY_NAME) ; fi
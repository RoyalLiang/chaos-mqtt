MAKEFILE_DIR:=$(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))
CWD=$(MAKEFILE_DIR)
GIT?=git
GO?=go
GORELEASER?=goreleaser

ARCH:=$(shell uname -s)
DIST:=$(MAKEFILE_DIR)/dist

PRODUCT=chaos

BASE_COMMIT:=$(shell git rev-list --first-parent HEAD | tail -n 1)

VERSION_MAJOR:=1
VERSION_MINOR:=0
VERSION_PATCH:=0
VERSION_SUFFIX:=$(shell git rev-list --count $(BASE_COMMIT)..HEAD)
VERSION:=$(VERSION_MAJOR).$(VERSION_MINOR).$(VERSION_PATCH)-$(VERSION_SUFFIX)

ifeq ($(OS),Windows_NT)
	PKG=$(PRODUCT).exe
else
	PKG=$(DIST)/$(PRODUCT).exe
endif

chaos: build windows darwin
	GOOS=linux CGO_ENABLED=0 $(GO) build -a -o $(DIST)/$(PRODUCT) -ldflags "-s -w -X main.Version=$(VERSION)" ./cmd/chaos

windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 $(GO) build -a -o $(PKG) -ldflags "-s -w -X main.Version=$(VERSION)" ./cmd/chaos

darwin:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 $(GO) build -a -o $(DIST)/$(PRODUCT)_darwin -ldflags "-s -w -X main.Version=$(VERSION)" ./cmd/chaos

build: clean
	@echo "Building $(PRODUCT), Version $(VERSION)"

clean:
	@for item in $(find $(DIST)/ -name "chaos*"); do	\
		$(RM) $$item; \
	done

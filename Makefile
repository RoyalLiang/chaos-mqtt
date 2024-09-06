MAKEFILE_DIR:=$(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))
CWD=$(MAKEFILE_DIR)
GIT?=git
GO?=go
GORELEASER?=goreleaser

ARCH:=$(shell uname -s)

PRODUCT=chaos

BASE_COMMIT:=$(shell git rev-list --first-parent HEAD | tail -n 1)

VERSION_MAJOR:=1
VERSION_MINOR:=0
VERSION_PATCH:=1
VERSION_SUFFIX:=$(shell git rev-list --count $(BASE_COMMIT)..HEAD)
VERSION:=$(VERSION_MAJOR).$(VERSION_MINOR).$(VERSION_PATCH)-$(VERSION_SUFFIX)

ifeq ($(OS),Windows_NT)
	EXTENSION=.exe
else
	EXTENSION=
endif

chaos: build
	CGO_ENABLED=0 $(GO) build -o $(PRODUCT)$(EXTENSION) -ldflags "-s -w -X main.Version=$(VERSION)" ./cmd/chaos

build: clean
	@echo "Building $(PRODUCT), Version $(VERSION)"

clean:
	@for item in $(find $(MAKEFILE_DIR)/ -name "chaos.exe"); do	\
		$(RM) $$item; \
	done

GO ?= go

OCIREPO ?= quay.io/aanm/junit2md
DOCKER ?= docker

all: tests junit2md

.PHONY: tests
tests:
	$(GO) test -cover ./...

junit2md: FORCE
	CGO_ENABLED=0 $(GO) build -o $@ ./cmd/...
	strip $@

.PHONY: image
image:
	$(DOCKER) build -f Dockerfile -t $(OCIREPO) .

.PHONY: install
install:
	CGO_ENABLED=0 $(GO) install ./cmd

clean:
	rm -f junit2md
FORCE:

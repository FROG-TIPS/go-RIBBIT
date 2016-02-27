GO := go
GOBUILD := $(GO) build -i 
TARGET_DIR := ./target

test:
	$(GO) test

distclean:
	rm -rf $(TARGET_DIR)

build:
	$(GOBUILD) -o $(TARGET_DIR)/cli frogsay/main.go

all: build

.PHONY: all build test distclean

TARGETS := $(shell find ./cmd -type d -maxdepth 1 -mindepth 1 -exec basename {} \;)
CMD_DIR := cmd
OUT_DIR := bin

all: build test

build: test
	@for target in $(TARGETS); do \
			go build -v -o $(OUT_DIR)/$${target} ./$(CMD_DIR)/$${target}; \
	done

test:
	@for target in $(TARGETS); do \
			go test -v ./$(CMD_DIR)/$${target}; \
	done

TARGETS := $(shell find ./cmd -type d -maxdepth 1 -mindepth 1 -exec basename {} \;)
CMD_DIR := cmd
OUT_DIR := bin

all: build

build:
	@for target in $(TARGETS); do                                        \
			go build -v -o $(OUT_DIR)/$${target} ./$(CMD_DIR)/$${target}; \
	done

.PHONY: all build run clean

BINARY_NAME = tetris-game
BUILD_DIR = build
HEIGHT ?= 20
WIDTH ?= 10

all: build

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go

run: build
	./$(BUILD_DIR)/$(BINARY_NAME) -height=$(HEIGHT) -width=$(WIDTH)

clean:
	rm -rf $(BUILD_DIR)

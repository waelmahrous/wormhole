BIN_DIR := bin

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/wormhole

SHELL := /bin/bash
.SILENT:

BIN_DIR := $(HOME)/.local/bin
ZSH_SYS_DIR := $(HOME)/.zsh/completions
BASH_SYS_DIR := $(HOME)/.bash_completion.d

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/wormhole

ZSH_SYS_DIR := $(HOME)/.zsh/completions

install-zsh-completion:
	mkdir -p $(ZSH_SYS_DIR)
	$(BIN_DIR)/wormhole completion zsh > $(ZSH_SYS_DIR)/_wormhole
	@echo "Zsh completion installed to $(ZSH_SYS_DIR)/_wormhole"

# COMBINED INSTALL
install: build install-zsh-completion
	@echo "wormhole installed with zsh completions."

SHELL := /bin/bash
.SILENT:

BIN_DIR := $(HOME)/.local/bin
ZSH_SYS_DIR := $(HOME)/.zsh/completions
BASH_SYS_DIR := $(HOME)/.bash_completion.d

build:
	@mkdir -p "$(BIN_DIR)"
	@go build -o "$(BIN_DIR)/wormhole"

install-zsh-completion: build
	@mkdir -p "$(ZSH_SYS_DIR)"
	@"$(BIN_DIR)/wormhole" completion zsh > "$(ZSH_SYS_DIR)/_wormhole"

install-bash-completion: build
	@mkdir -p "$(BASH_SYS_DIR)"
	@"$(BIN_DIR)/wormhole" completion bash > "$(BASH_SYS_DIR)/wormhole"

install: build install-zsh-completion install-bash-completion
	@echo "wormhole installed"

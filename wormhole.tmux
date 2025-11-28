#!/usr/bin/env bash

CURRENT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Defaults
DEFAULT_KEY="O"
DEFAULT_KEY_CURRENT="C-o"

# Read user-configurable options (if set)
WORMHOLE_KEY="$(tmux show-option -gqv "@wormhole_key")"
WORMHOLE_KEY_CURRENT="$(tmux show-option -gqv "@wormhole_key_current")"

# Fall back to defaults if empty
[ -z "$WORMHOLE_KEY" ] && WORMHOLE_KEY="$DEFAULT_KEY"
[ -z "$WORMHOLE_KEY_CURRENT" ] && WORMHOLE_KEY_CURRENT="$DEFAULT_KEY_CURRENT"

# Install plugin
tmux run-shell "${CURRENT_DIR}/wormhole.sh install"

# Choose an open pane
tmux bind-key "$WORMHOLE_KEY" choose-tree -Z \
	"run-shell \"${CURRENT_DIR}/wormhole.sh open '%%'\""

# Open in current pane
tmux bind-key "$WORMHOLE_KEY_CURRENT" \
	"run-shell \"${CURRENT_DIR}/wormhole.sh open\""

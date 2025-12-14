#!/usr/bin/env bash

CURRENT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Defaults
DEFAULT_KEY="O"
DEFAULT_KEY_CURRENT="C-o"
DEFAULT_DEBUG="false"
DEFAULT_KEY_JUMP="j"

# Read user-configurable options (if set)
WORMHOLE_KEY="$(tmux show-option -gqv "@wormhole_key")"
WORMHOLE_KEY_CURRENT="$(tmux show-option -gqv "@wormhole_key_current")"
WORMHOLE_DEBUG="$(tmux show-option -gqv "@wormhole_debug")"
WORMHOLE_KEY_JUMP="$(tmux show-option -gqv "@wormhole_key_jump")"

# Fall back to defaults if empty
[ -z "$WORMHOLE_KEY" ] && WORMHOLE_KEY="$DEFAULT_KEY"
[ -z "$WORMHOLE_KEY_CURRENT" ] && WORMHOLE_KEY_CURRENT="$DEFAULT_KEY_CURRENT"
[ -z "$WORMHOLE_DEBUG" ] && WORMHOLE_DEBUG="$DEFAULT_DEBUG"
[ -z "$WORMHOLE_KEY_JUMP" ] && WORMHOLE_KEY_JUMP="$DEFAULT_KEY_JUMP"

# Convert true/false into CLI flag
[ "$WORMHOLE_DEBUG" = "true" ] && DEBUG_FLAG="--debug" || DEBUG_FLAG=""

# Install plugin (pass flag into wormhole.sh, not the raw true/false)
tmux run-shell "${CURRENT_DIR}/wormhole.sh install ${DEBUG_FLAG}"

# Choose an open pane
tmux bind-key "$WORMHOLE_KEY" choose-tree -Z \
  "run-shell \"${CURRENT_DIR}/wormhole.sh open '%%' ${DEBUG_FLAG}\""

# Open in current pane
tmux bind-key "$WORMHOLE_KEY_CURRENT" \
  "run-shell \"${CURRENT_DIR}/wormhole.sh open ${DEBUG_FLAG}\""

# Open new window at wormhole
tmux bind-key "$WORMHOLE_KEY_JUMP" \
  "run-shell \"${CURRENT_DIR}/wormhole.sh jump\""

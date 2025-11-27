#!/usr/bin/env bash

CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

DEFAULT_KEY=W
DEFAULT_KEY_CURRENT=C-w

# Install plugin
tmux run-shell "${CURRENT_DIR}/wormhole.sh install"

# Choose an open pane
tmux bind-key $DEFAULT_KEY choose-tree -Z -F \
  "#{window_index}.#{pane_index}" \
  "run-shell \"${CURRENT_DIR}/wormhole.sh open '%%'\""

# Open in current pane
tmux bind-key $DEFAULT_KEY_CURRENT \
  "run-shell \"${CURRENT_DIR}/wormhole.sh open\""

#!/usr/bin/env bash

CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

DEFAULT_KEY=W
DEFAULT_KEY_CURRENT=C-w

# Choose an open pane
tmux bind-key $DEFAULT_KEY choose-tree -Z -F \
  "#{window_index}.#{pane_index}" \
  "run-shell \"${CURRENT_DIR}/scripts/wormhole.sh open '%%'\""

# Open in current pane
tmux bind-key $DEFAULT_KEY_CURRENT \
  "run-shell \"${CURRENT_DIR}/scripts/wormhole.sh open\""

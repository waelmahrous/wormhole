#!/usr/bin/env bash

CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

DEFAULT_KEY=W

tmux bind-key $DEFAULT_KEY choose-tree -Z -F \
  "#{window_index}.#{pane_index}" \
  "run-shell \"${CURRENT_DIR}/scripts/wormhole.sh open '%%'\""

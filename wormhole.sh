#!/usr/bin/env bash

cmd="$1"
pane="$2"

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
WORMHOLE_BIN="$HOME/.local/bin/wormhole"

case "$cmd" in
	install)
		( cd $SCRIPT_DIR && make install )
		;;
	open)
		dir=$(tmux display-message -p -t "$pane" "#{pane_current_path}")
		out="$($WORMHOLE_BIN open --destination "$dir" 2>&1)"
		tmux display-message "$out"
		exit 0
		;;
	*)
		tmux display-message "wormhole.sh: invalid command '$cmd'"
		exit 1
		;;
esac

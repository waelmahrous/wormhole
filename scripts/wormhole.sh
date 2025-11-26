#!/usr/bin/env bash

cmd="$1"
pane="$2"

case "$cmd" in
	open)
		dir=$(tmux display-message -p -t "$pane" "#{pane_current_path}")
		out="$(wormhole open --destination "$dir" 2>&1)"
		tmux display-message "$out"
		exit 0
		;;
	*)
		tmux display-message "wormhole.sh: invalid command '$cmd'"
		exit 1
		;;
esac

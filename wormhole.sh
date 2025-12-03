#!/usr/bin/env bash

cmd="$1"
pane="$2"

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
WORMHOLE_BIN="$HOME/.local/bin/wormhole"

os=$(uname -s)
arch=$(uname -m)

case "$cmd" in
	install)
		case "$os" in
			Linux)  goos="linux"  ;;
			Darwin) goos="darwin" ;;
			*) echo "wormhole: unsupported OS: $os" >&2; exit 1 ;;
		esac

		case "$arch" in
			x86_64|amd64)  goarch="amd64" ;;
			arm64|aarch64) goarch="arm64" ;;
			*) echo "wormhole: unsupported arch: $arch" >&2; exit 1 ;;
		esac

		tmpdir="$(mktemp -d)"
		trap 'rm -rf "$tmpdir"' EXIT

		asset="wormhole_${goos}_${goarch}.tar.gz"
		url="https://github.com/waelmahrous/wormhole/releases/latest/download/${asset}"

		echo "wormhole: downloading ${asset} from ${url}..."
		if curl -fL -o "${tmpdir}/${asset}" "${url}"; then
			echo "wormhole: download ok, installing..."

			# Extract archive
			tar -xzf "${tmpdir}/${asset}" -C "${tmpdir}"

			# Binary inside archive is built as: wormhole-<goos>-<goarch>
			src_bin="${tmpdir}/wormhole-${goos}-${goarch}"

			if [ ! -x "${src_bin}" ]; then
				echo "wormhole: expected binary not found in archive: ${src_bin}" >&2
				exit 1
			fi

			install -d "$(dirname "${WORMHOLE_BIN}")"
			install "${src_bin}" "${WORMHOLE_BIN}"

			echo "wormhole: installed to ${WORMHOLE_BIN}"
		else
			echo "wormhole: failed to download prebuilt binary, falling back to build from source..."
			(
				cd "${SCRIPT_DIR}" || exit 1
				make install
			)
		fi
		;;
	open)
		dir=$(tmux display-message -p -t "$pane" "#{pane_current_path}")
		out="$("$WORMHOLE_BIN" open --destination "$dir" 2>&1)"
		tmux display-message "$out"
		exit 0
		;;
	*)
		tmux display-message "wormhole.sh: invalid command '$cmd'"
		exit 1
		;;
esac

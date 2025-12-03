#!/usr/bin/env bash

cmd="$1"
pane="$2"
debug=false

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
WORMHOLE_BIN="$HOME/.local/bin/wormhole"

os=$(uname -s)
arch=$(uname -m)

for arg in "$@"; do
    case "$arg" in
        --debug)
            debug=true
            ;;
    esac
done

print() {
    if [ "$debug" = "false" ]; then
        return
    fi

    echo "$@"
}

case "$cmd" in
	install)
		case "$os" in
			Linux)  goos="linux"  ;;
			Darwin) goos="darwin" ;;
			*) print "wormhole: unsupported OS: $os" >&2; exit 1 ;;
		esac

		case "$arch" in
			x86_64|amd64)  goarch="amd64" ;;
			arm64|aarch64) goarch="arm64" ;;
			*) print "wormhole: unsupported arch: $arch" >&2; exit 1 ;;
		esac

		tmpdir="$(mktemp -d)"
		trap 'rm -rf "$tmpdir"' EXIT

		asset="wormhole_${goos}_${goarch}.tar.gz"
		url="https://github.com/waelmahrous/wormhole/releases/latest/download/${asset}"

		print "wormhole: downloading ${asset} from ${url}..."
		if curl -fL -o "${tmpdir}/${asset}" "${url}"; then
			print "wormhole: download ok, installing..."

			# Extract archive
			tar -xzf "${tmpdir}/${asset}" -C "${tmpdir}"

			# Binary inside archive is built as: wormhole-<goos>-<goarch>
			src_bin="${tmpdir}/wormhole-${goos}-${goarch}"

			if [ ! -x "${src_bin}" ]; then
				print "wormhole: expected binary not found in archive: ${src_bin}" >&2
				exit 1
			fi

			install -d "$(dirname "${WORMHOLE_BIN}")"
			install "${src_bin}" "${WORMHOLE_BIN}"

			print "wormhole: installed to ${WORMHOLE_BIN}"
		else
			print "wormhole: failed to download prebuilt binary, falling back to build from source..."
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

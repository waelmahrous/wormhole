#!/usr/bin/env bash

# Where we store the current wormhole destination
STATE_DIR="${XDG_STATE_HOME:-$HOME/.local/state}/wormhole"
STATE_FILE="$STATE_DIR/target"

usage() {
  echo "Usage:"
  echo "  wormhole open           # set current directory as wormhole destination"
  echo "  wormhole send FILE...   # send one or more files/dirs to destination"
  echo "  wormhole copy FILE...   # copy one or more files/dirs to destination"
  echo "  wormhole pull           # pull all files (copy) near wormhole"
  echo "  wormhole jump           # jump inside of wormhole"
  echo "  wormhole status         # show current destination"
  echo "  wormhole close          # remove current destination"
}

cmd="$1"
shift || true

case "$cmd" in
  open)
    dest="$(pwd)"
    mkdir -p "$STATE_DIR"
    printf '%s\n' "$dest" > "$STATE_FILE"
    echo "Wormhole destination set to:"
    echo "  $dest"
    ;;

  send | copy | pull)
    if [ ! -f "$STATE_FILE" ]; then
      echo "Error: no wormhole destination set. Run 'wormhole open' first." >&2
      exit 1
    fi

    dest="$(cat "$STATE_FILE")"

    if [ "$#" -eq 0 ] && [ "$cmd" != pull ]; then
      echo "Error: no files to send." >&2
      echo "Usage: wormhole send FILE..." >&2
      exit 1
    fi

    # Copy all args into the destination, keeping basenames
    if [ "$cmd" == pull ]; then
      cp -R -- "$dest"/* "$(pwd)"
    elif [ "$cmd" == copy ]; then
      cp -R -- "$@" "$dest"/
    else
      mv -- "$@" "$dest"/
    fi

    [ "$cmd" = pull ] && count="all" || count="$#"

    echo "$cmd $count item(s) to:"
    [ "$cmd" = pull ] && echo $(pwd) || echo "$dest"

    ;;

  jump)
    if [ ! -f "$STATE_FILE" ]; then
      echo "No wormhole destination set."
      exit 1
    fi

    dest="$(cat "$STATE_FILE")"

    # Change to the destination
    cd "$dest" || exit 1

    echo "Jumping to $dest"

    # Replace the script process with a new shell
    exec "$SHELL"
    ;;

  status)
    if [ -f "$STATE_FILE" ]; then
      echo "Current wormhole destination:"
      cat "$STATE_FILE"
    else
      echo "No wormhole destination set."
    fi
    ;;

  close)
    rm -f "$STATE_FILE"
    echo "Wormhole destination closed."
    ;;

  ""|help|-h|--help)
    usage
    ;;

  *)
    echo "Unknown command: $cmd" >&2
    usage
    exit 1
    ;;
esac

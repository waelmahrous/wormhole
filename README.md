# wormhole

### Teleport files between terminals without thinking about paths.

**wormhole** is a tiny shell utility that lets you “open” a destination in one terminal, then instantly send or copy to it from anywhere else. No more typing long paths, no more hunting for the right directory. If you already have a terminal open where you want things to go, that location becomes your wormhole.

Might turn this into a tmux plugin, there is minimal support for that atm.

---

- [Features](#features)
- [Installation](#installation)
  - [Go](#option-1-install-via-go-install)
  - [Build from source](#option-2-build-from-source)
  - [tmux integration](#tmux-TPM)
- [Example Workflow](#example-workflow)
  - [shell](#shell-a-navigate-where-you-want-files-to-land)
  - [tmux usage](#tmux-usage)

## Features

- **Open a wormhole** in the current directory
- **Send** files or directories into the wormhole
- **Copy** files without deleting the originals
- **Check status** 
- **tmux integration**

## Installation

Make sure you have **Go** installed.

### Option 1: Install via go install

```Bash
go install github.com/waelmahrous/wormhole@latest
```

Make sure your Go bin directory is in your PATH:

```Bash
export PATH="$HOME/go/bin:$PATH"
```

Add that line to your shell config (~/.zshrc, ~/.bashrc, etc.) to make it permanent.

### Option 2: Build from source

Clone the repository and run:

```Bash
make install
```

This builds the application and installs shell autocompletion.

### tmux TPM

Add this line to your .tmux.conf to install the plugin with TPM:

```Bash
set -g @plugin 'waelmahrous/wormhole'
```

## Example Workflow

### Shell A: navigate where you want files to land
```Bash
cd ~/projects/myapp/build
wormhole open

```
### Shell B: anywhere else
```Bash
wormhole send file1.txt assets/*.png
```

### tmux usage

wormhole comes with two default keybindings which you can alter:

```Bash
set -g @wormhole_key "O"            # open tree selector
set -g @wormhole_key_current "C-o"  # choose current pane as destination
```

```Bash
# Opens a tree selector. Pick a pane and a wormhole will open using that pane’s working directory.
Prefix-O
```

```Bash
# Instantly opens a wormhole at the current pane’s working directory (no picker).
Prefix-Ctrl-o
```

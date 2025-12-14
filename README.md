# wormhole

### Teleport files between terminals without thinking about paths.

**wormhole** is a tiny shell utility that lets you “open” a destination in one terminal, then instantly send or copy to it from anywhere else. No more typing long paths, no more hunting for the right directory. If you already have a terminal open where you want things to go, that location becomes your wormhole.

Might turn this into a tmux plugin.

---

- [Features](#features)
- [Installation](#installation)
  - [TPM](#install-via-tmux-tpm-recommended)
  - [Go](#install-via-go-install)
  - [Build from source](#build-from-source)
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

The recommended way to install **wormhole** is through the tmux plugin manager (TPM). Make sure Go is installed.

### **Install via tmux TPM (recommended)**

Add the following to your `.tmux.conf`:

```bash
set -g @plugin 'waelmahrous/wormhole'
```

Then press:

```Bash
prefix + I
```

TPM will automatically download the plugin, build the wormhole binary, and install completions.

### Install via go install

```Bash
go install github.com/waelmahrous/wormhole@latest
```

Make sure your Go bin directory is in your PATH:

```Bash
export PATH="$HOME/go/bin:$PATH"
```

Add that line to your shell config (~/.zshrc, ~/.bashrc, etc.) to make it permanent.

### Build from source

Clone the repository and run:

```Bash
make install
```

This builds the application and installs shell autocompletion.

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

wormhole comes with the following default keybindings which you can alter:

```Bash
set -g @wormhole_key "O"            # open tree selector
set -g @wormhole_key_current "C-o"  # choose current pane as destination
set -g @wormhole_debug "false"      # show install logs
set -g @wormhole_key_jump "j"       # Open new window at wormhole
```

```Bash
# Opens a tree selector. Pick a pane and a wormhole will open using that pane’s working directory.
prefix-O
```

```Bash
# Instantly opens a wormhole at the current pane’s working directory (no picker).
prefix-Ctrl-o
```

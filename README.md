# wormhole

### Teleport files between terminals without thinking about paths.

**wormhole** is a tiny shell utility that lets you “open” a destination in one terminal, then instantly send or copy to it from anywhere else. No more typing long paths, no more hunting for the right directory. If you already have a terminal open where you want things to go, that location becomes your wormhole.

---

## Features

- **Open a wormhole** in the current directory
- **Send** files or directories into the wormhole
- **Copy** files without deleting the originals
- **Check status** 

## Installation


### Option 1: Install via go install

```Bash
go install github.com/waelmahrous/wormhole/cmd/wormhole@latest
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


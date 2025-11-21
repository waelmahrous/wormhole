# wormhole

### Teleport files between terminals without thinking about paths.

**wormhole** is a tiny shell utility that lets you “open” a destination in one terminal, then instantly send, copy, pull, or jump to it from anywhere else. No more typing long paths, no more hunting for the right directory. If you already have a terminal open where you want things to go, that location becomes your wormhole.

---

## Features

- **Open a wormhole** in the current directory
- **Send** files or directories into the wormhole
- **Copy** files without deleting the originals
- **Pull** everything from the wormhole back to your current location
- **Jump** into the wormhole destination (spawns a new shell session there)
- **Check status** or **close** the wormhole

## Installation

Run `make build` to build the application.

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


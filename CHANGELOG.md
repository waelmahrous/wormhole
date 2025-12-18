## v0.8.0 (2025-12-18)

### Feat

- **send**: add force mode
- add a safezone
- **root**: set configurable wormhole id
- **args**: add args struct for wormhole
- **tmux**: add keybinding to open new window at wormhole
- **send**: use GetDestination and store transfer records in database

### Fix

- **root**: dont return on status

### Refactor

- cleanup db interaction
- add withDB to wormhole methods
- **safezone**: small cleanup
- **root**: use switch case for flags
- convert static state functions to wormhole methods

## v0.7.2 (2025-12-12)

### Fix

- **root**: create state file on any command

### Refactor

- **internal**: add withDB helper and introduce transfer record

## v0.7.1 (2025-12-10)

### Fix

- **send**: return error when sending directories or overwriting files

## v0.7.0 (2025-12-10)

### BREAKING CHANGE

- existing .wormhole.json state will no longer be read.

### Feat

- add to path on install
- silence tmux output if not debugging

### Refactor

- change wormhole backend from json to storm
- remove unnecessary internal logging
- **send**: add tests and send by copy first

## v0.6.0 (2025-12-03)

### Feat

- default to install binaries
- build with make
- upload binary artifacts

## v0.5.1 (2025-12-03)

### Fix

- **send**: use correct path to state directory

## v0.5.0 (2025-12-02)

### Feat

- **tmux**: change default keybindings (#2)
- **tmux**: allow user to change default keybindings

### Fix

- **tmux**: remove choose-tree formatting

### Refactor

- migrate to directory-based state management and unify root command behavior
- **root**: expose root command
- **open**: cleanup
- **send**: cleanup

## v0.4.0 (2025-11-27)

### Feat

- add bash completion
- **tmux**: integrate with TPM

## v0.3.0 (2025-11-26)

### Feat

- **tmux**: add binding to open wormhole in current pane
- **tmux**: minimal integration
- **open**: add custom destination flag

### Refactor

- change state file to json format

## v0.2.0 (2025-11-25)

### Feat

- **autocompletion**: add completion step do installation script
- **root**: add flag to see wormhole status
- **send**: add copy flag to send command
- **send**: add command to send files through wormhole
- **open**: add command to create wormhole at working directory

### Fix

- **send**: handle directories

### Refactor

- **root**: turn verbose flag into silent flag

## v0.1.0 (2025-11-25)

### Feat

- **root**: add verbose flag for debug logging
- **open**: add open command
- basic build

package internal

import (
	"errors"
	"os"
	"path/filepath"
)

type WormholeState struct {
	Destination string `json:"destination"`
}

const defaultStateFile = ".wormhole.json"

func resolveStatePath(path string) (string, error) {
	if path != "" {
		if _, err := os.ReadDir(path); err != nil {
			return "", err
		}
		return filepath.Join(path, defaultStateFile), nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, defaultStateFile), nil
}

func SaveWormholeState(path string, state WormholeState) error {
	statePath, err := resolveStatePath(path)
	if err != nil {
		return err
	}
	return SaveJSON(statePath, state)
}

func LoadWormholeState(path string) (WormholeState, error) {
	statePath, err := resolveStatePath(path)
	if err != nil {
		return WormholeState{}, err
	}

	var state WormholeState
	err = LoadJSON(statePath, &state)

	if os.IsNotExist(err) {
		return WormholeState{}, nil
	}

	return state, err
}

func UpdateDestination(path, dest string) error {
	state, err := LoadWormholeState(path)
	if err != nil {
		return err
	}

	if _, err := os.ReadDir(dest); err != nil {
		return err
	}

	state.Destination = dest
	return SaveWormholeState(path, state)
}

func GetDestination(path string) (string, error) {
	state, err := LoadWormholeState(path)
	if err != nil {
		return "", err
	}
	if state.Destination == "" {
		return "", errors.New("no wormhole destination set")
	}
	return state.Destination, nil
}

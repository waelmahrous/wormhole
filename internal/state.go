package internal

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/asdine/storm/v3"
	"github.com/otiai10/copy"
)

const (
	StoreName          = ".wormhole.db"
	DefaultID          = "0"
	DefaultDestination = ""
)

type Wormhole struct {
	ID          string `storm:"id"`
	Destination string
}

func SetDestination(path string, target string) (Wormhole, error) {
	if db, err := storm.Open(filepath.Join(path, StoreName)); err != nil {
		return Wormhole{}, err
	} else {
		defer db.Close()

		var wormhole Wormhole

		if err := db.One("ID", DefaultID, &wormhole); err != nil {
			return Wormhole{}, err
		} else {
			wormhole.Destination = target
			if err := db.Save(&wormhole); err != nil {
				return Wormhole{}, err
			}

			return wormhole, nil
		}
	}
}

func GetDestination(path string) (string, error) {
	if db, err := storm.Open(filepath.Join(path, StoreName)); err != nil {
		return "", err
	} else {
		defer db.Close()

		var wormhole Wormhole

		if err := db.One("ID", DefaultID, &wormhole); err != nil {
			return "", err
		} else {
			if wormhole.Destination == "" {
				return "", errors.New("no wormhole open")
			}

			return wormhole.Destination, nil
		}
	}
}

func InitWormholeStore(path string) error {
	if path == "" {
		return errors.New("empty state directory")
	}

	if _, err := os.Stat(filepath.Join(path, StoreName)); err == nil {
		return nil
	}

	if db, err := storm.Open(filepath.Join(path, StoreName)); err != nil {
		return err
	} else {
		defer db.Close()

		wormhole := Wormhole{
			ID:          DefaultID,
			Destination: DefaultDestination,
		}

		if err := db.Save(&wormhole); err != nil {
			return err
		}

		return nil
	}
}

func Transfer(src []string, dst string, copyMode bool) ([]string, error) {
	if len(src) < 1 {
		return nil, errors.New("no files to send")
	}

	output := []string{}

	for _, v := range src {
		if _, err := os.ReadDir(v); err == nil {
			return output, errors.New("this is a directory")
		}

		filePath := filepath.Join(filepath.Join(dst, filepath.Base(v)))

		if _, err := os.Stat(filePath); err == nil {
			return output, errors.New("file already exists in target directory")
		}

		if err := copy.Copy(v, filePath); err != nil {
			return output, err
		}

		output = append(output, filePath)

		if copyMode {
			continue
		}

		if err := os.Remove(v); err != nil {
			return output, err
		}
	}

	return output, nil
}

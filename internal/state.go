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

type TransferRecord struct {
	ID          int `storm:"id,increment"`
	Source      []string
	Destination string
	Copy        bool
	StateDir    string
}

type Operation func(*storm.DB) error

func withDB(path string, op Operation) error {
	if db, err := storm.Open(filepath.Join(path, StoreName)); err != nil {
		return err
	} else {
		defer db.Close()
		return op(db)
	}
}

func SetDestination(path string, target string) (Wormhole, error) {
	var wormhole Wormhole

	err := withDB(path, func(db *storm.DB) error {
		if err := db.One("ID", DefaultID, &wormhole); err != nil {
			return err
		}

		wormhole.Destination = target
		return db.Save(&wormhole)
	})

	return wormhole, err
}

func GetDestination(path string) (string, error) {
	var err error
	var destination string

	err = withDB(path, func(db *storm.DB) error {
		var wormhole Wormhole

		if err := db.One("ID", DefaultID, &wormhole); err != nil {
			return err
		} else {
			if wormhole.Destination == "" {
				return errors.New("no wormhole open")
			}

			destination = wormhole.Destination
			return nil
		}
	})

	return destination, err
}

func InitWormholeStore(path string) error {
	if path == "" {
		return errors.New("empty state directory")
	}

	wormhole := Wormhole{
		ID:          DefaultID,
		Destination: DefaultDestination,
	}

	return withDB(path, func(db *storm.DB) error {
		if err := db.One("ID", DefaultID, &wormhole); err != nil {
			return db.Save(&wormhole)
		}

		return nil
	})
}

func Transfer(record TransferRecord) ([]string, error) {
	if len(record.Source) < 1 {
		return nil, errors.New("no files to send")
	}

	output := []string{}

	for _, v := range record.Source {
		if _, err := os.ReadDir(v); err == nil {
			return output, errors.New("this is a directory")
		}

		filePath := filepath.Join(filepath.Join(record.Destination, filepath.Base(v)))

		if _, err := os.Stat(filePath); err == nil {
			return output, errors.New("file already exists in target directory")
		}

		if err := copy.Copy(v, filePath); err != nil {
			return output, err
		}

		output = append(output, filePath)

		if record.Copy {
			continue
		}

		if err := os.Remove(v); err != nil {
			return output, err
		}
	}

	return output, nil
}

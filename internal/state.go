package internal

import (
	"errors"
	"fmt"
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
	Copy        bool
	StateDir    string
	Destination string
	WorkDir     string
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
	if len(record.Source) == 0 {
		return nil, errors.New("no files to send")
	}

	destination, err := GetDestination(record.StateDir)
	if err != nil {
		return nil, err
	}

	var output []string

	for _, src := range record.Source {
		if _, err := os.ReadDir(src); err == nil {
			return output, errors.New("this is a directory")
		}

		filePath := filepath.Join(destination, filepath.Base(src))

		if _, err := os.Stat(filePath); err == nil {
			return output, errors.New("file already exists in target directory")
		}

		if err := copy.Copy(src, filePath); err != nil {
			return output, err
		}

		output = append(output, filePath)

		if !record.Copy {
			if err := os.Remove(src); err != nil {
				return output, err
			}
		}
	}

	return output, withDB(record.StateDir, func(db *storm.DB) error {
		record.Destination = destination
		if wd, err := os.Getwd(); err != nil {
			return fmt.Errorf("could not establish working directory, %v", err)
		} else {
			record.WorkDir = wd
		}

		return db.Save(&record)
	})
}

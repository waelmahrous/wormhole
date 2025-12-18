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
	DefaultDestination = ""
	DefaultSafeZone    = ".wormhole"
)

type Wormhole struct {
	ID          string `storm:"id"`
	Destination string
	StateDir    string
	Args        WormholeArgs
}

type WormholeArgs struct {
}

type TransferRecord struct {
	ID          int    `storm:"id,increment"`
	WormholeID  string `storm:"index"`
	Source      []string
	Copy        bool
	Destination string
	WorkDir     string
	Force       bool
}

type Operation func(*storm.DB) error

func (w *Wormhole) withDB(op Operation) error {
	if db, err := storm.Open(filepath.Join(w.StateDir, StoreName)); err != nil {
		return err
	} else {
		defer db.Close()

		if err := db.One("ID", w.ID, w); err != nil {
			return err
		}

		return op(db)
	}
}

func (w *Wormhole) SetDestination(target string) error {
	return w.withDB(func(db *storm.DB) error {
		w.Destination = target
		return db.Save(w)
	})
}

func (w *Wormhole) GetDestination() (string, error) {
	var destination string

	return destination, w.withDB(func(db *storm.DB) error {
		if w.Destination == "" {
			return errors.New("no wormhole open")
		}

		destination = w.Destination
		return nil
	})
}

func (w *Wormhole) SetArgs(a WormholeArgs) error {
	return w.withDB(func(db *storm.DB) error {
		w.Args = a
		return db.Save(w)
	})
}

func (w *Wormhole) InitWormholeStore() error {
	if w.StateDir == "" {
		return errors.New("empty state directory")
	}

	if db, err := storm.Open(filepath.Join(w.StateDir, StoreName)); err != nil {
		return err
	} else {
		defer db.Close()

		if err := db.One("ID", w.ID, w); err != nil {
			return db.Save(w)
		}

		return nil
	}
}

func (w *Wormhole) Transfer(record TransferRecord) ([]string, error) {
	if len(record.Source) == 0 {
		return nil, errors.New("no files to send")
	}

	destination, err := w.GetDestination()
	if err != nil {
		return nil, err
	}

	var output []string

	for _, src := range record.Source {
		if _, err := os.ReadDir(src); err == nil {
			return output, errors.New("this is a directory")
		}

		filePath := filepath.Join(destination, filepath.Base(src))

		if _, err := os.Stat(filePath); err == nil && record.Force == false {
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

	return output, w.withDB(func(db *storm.DB) error {
		record.Destination = destination
		if wd, err := os.Getwd(); err != nil {
			return fmt.Errorf("could not establish working directory, %v", err)
		} else {
			record.WorkDir = wd
		}

		return db.Save(&record)
	})
}

package internal

import (
	"encoding/json"
	"log"
	"os"
)

func SaveJSON(path string, v any) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

func LoadJSON(path string, v any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func Fatalf(format string, a ...any) {
	log.Printf(format, a...)
	os.Exit(1)
}

package fs

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func appendToJSON[T any](file string, v T) error {
	if err := initJSON(file); err != nil {
		return err
	}

	items, err := decodeJSON[T](file)
	if err != nil {
		return err
	}

	items = append(items, v)

	return writeJSON(file, items)
}

func decodeJSON[T any](file string) ([]T, error) {
	if err := initJSON(file); err != nil {
		return nil, err
	}

	f, err := os.OpenFile(file, os.O_RDONLY, 0640)
	if err != nil {
		return nil, err
	}

	var items []T

	if err := json.NewDecoder(f).Decode(&items); err != nil {
		return nil, fmt.Errorf("error reading from %q: %v", file, err)
	}

	return items, nil
}

func writeJSON(file string, v any) error {
	if err := initJSON(file); err != nil {
		return err
	}

	f, err := os.CreateTemp("", "")
	if err != nil {
		return fmt.Errorf("error creating temp file: %v", err)
	}

	var enc = json.NewEncoder(f)
	enc.SetIndent("", "  ")

	if err := enc.Encode(v); err != nil {
		return fmt.Errorf("error writing to %q: %v", file, err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("error writing to %q: %v", file, err)
	}

	if err := os.Rename(f.Name(), file); err != nil {
		return fmt.Errorf("error renaming temp file to %q: %v", file, err)
	}

	return nil
}

func initJSON(file string) error {
	if _, err := os.Stat(file); err == nil {
		return nil
	} else if !errors.Is(err, os.ErrNotExist) {
		return err
	}

	var dir = filepath.Dir(file)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("error creating %q: %w", dir, err)
	}

	f, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("error creating %q: %w", file, err)
	}

	if err := json.NewEncoder(f).Encode([]any{}); err != nil {
		return fmt.Errorf("error writing to %q: %w", file, err)
	}

	return f.Close()
}

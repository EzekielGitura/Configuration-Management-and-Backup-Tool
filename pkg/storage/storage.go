package storage

import (
    "fmt"
    "os"
)

type Storage struct {
    Path string
}

func NewStorage(path string) *Storage {
    return &Storage{
        Path: path,
    }
}

func (s *Storage) EnsurePathExists() error {
    // Ensure the storage path exists
    if _, err := os.Stat(s.Path); os.IsNotExist(err) {
        err := os.MkdirAll(s.Path, 0755)
        if err != nil {
            return fmt.Errorf("failed to create storage path: %v", err)
        }
    }
    return nil
}
package storage

import (
    "io/ioutil"
    "os"
    "path/filepath"
    "testing"
)

func TestNewStorage(t *testing.T) {
    // Create a temporary directory for testing
    tempDir, err := ioutil.TempDir("", "storage_test")
    if err != nil {
        t.Fatalf("Failed to create temporary directory: %v", err)
    }
    defer os.RemoveAll(tempDir)

    // Initialize storage client
    storageClient := NewStorage(tempDir)

    // Verify the storage path exists
    if _, err := os.Stat(storageClient.Path); os.IsNotExist(err) {
        t.Fatalf("Storage path does not exist: %v", err)
    }
}

func TestEnsurePathExists(t *testing.T) {
    // Create a temporary directory for testing
    tempDir, err := ioutil.TempDir("", "storage_test")
    if err != nil {
        t.Fatalf("Failed to create temporary directory: %v", err)
    }
    defer os.RemoveAll(tempDir)

    // Remove the directory to simulate non-existent path
    os.RemoveAll(tempDir)

    // Initialize storage client
    storageClient := NewStorage(tempDir)

    // Ensure the path exists
    err = storageClient.EnsurePathExists()
    if err != nil {
        t.Fatalf("Failed to ensure path exists: %v", err)
    }

    // Verify the storage path exists
    if _, err := os.Stat(storageClient.Path); os.IsNotExist(err) {
        t.Fatalf("Storage path does not exist: %v", err)
    }
}
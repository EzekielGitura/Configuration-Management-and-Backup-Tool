package backup

import (
    "io/ioutil"
    "os"
    "path/filepath"
    "testing"

    "networking-tool/pkg/storage"
)

func TestCreateBackup(t *testing.T) {
    // Create a temporary directory for testing
    tempDir, err := ioutil.TempDir("", "backup_test")
    if err != nil {
        t.Fatalf("Failed to create temporary directory: %v", err)
    }
    defer os.RemoveAll(tempDir)

    // Initialize storage client
    storageClient := storage.NewStorage(tempDir)

    // Network info for testing
    networkInfo := `{"status": "active", "devices": ["eth0", "eth1"]}`

    // Create a backup
    err = CreateBackup(storageClient, networkInfo)
    if err != nil {
        t.Fatalf("Failed to create backup: %v", err)
    }

    // Verify the backup file
    files, err := ioutil.ReadDir(tempDir)
    if err != nil {
        t.Fatalf("Failed to read directory: %v", err)
    }

    if len(files) != 1 {
        t.Fatalf("Expected 1 backup file, got %d", len(files))
    }

    filePath := filepath.Join(tempDir, files[0].Name())
    data, err := ioutil.ReadFile(filePath)
    if err != nil {
        t.Fatalf("Failed to read backup file: %v", err)
    }

    if string(data) != networkInfo {
        t.Fatalf("Backup file content mismatch: got %s, want %s", string(data), networkInfo)
    }
}
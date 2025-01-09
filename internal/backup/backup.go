package backup

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "time"

    "networking-tool/pkg/storage"
)

func CreateBackup(storageClient *storage.Storage, networkInfo string) error {
    // Create a backup file name with a timestamp
    fileName := fmt.Sprintf("backup_%s.json", time.Now().Format("20060102150405"))
    filePath := filepath.Join(storageClient.Path, fileName)

    // Write network info to the backup file
    err := ioutil.WriteFile(filePath, []byte(networkInfo), 0644)
    if err != nil {
        return fmt.Errorf("failed to write backup file: %v", err)
    }

    return nil
}
package main

import (
    "fmt"
    "log"
    "os"

    "networking-tool/internal/config"
    "networking-tool/internal/backup"
    "networking-tool/internal/network"
    "networking-tool/internal/utils"
    "networking-tool/pkg/api"
    "networking-tool/pkg/storage"
)

func main() {
    // Load configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }

    // Initialize API client
    apiClient := api.NewClient(cfg.API.URL, cfg.API.Token)

    // Initialize storage
    storageClient := storage.NewStorage(cfg.Storage.Path)

    // Example network operation
    networkInfo, err := network.GetNetworkInfo(apiClient)
    if err != nil {
        log.Fatalf("Failed to get network info: %v", err)
    }
    fmt.Println("Network Info:", networkInfo)

    // Example backup operation
    err = backup.CreateBackup(storageClient, networkInfo)
    if err != nil {
        log.Fatalf("Failed to create backup: %v", err)
    }
    fmt.Println("Backup created successfully")

    // Example utility operation
    utils.PrintVersion()
}
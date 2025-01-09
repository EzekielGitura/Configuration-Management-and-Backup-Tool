package config

import (
    "encoding/json"
    "os"
    "testing"
)

func TestLoadConfig(t *testing.T) {
    // Create a temporary config file for testing
    configData := `{
        "api": {
            "url": "https://api.example.com",
            "token": "testtoken"
        },
        "storage": {
            "path": "/tmp/backups"
        }
    }`
    configFile := "config_test.json"
    err := ioutil.WriteFile(configFile, []byte(configData), 0644)
    if err != nil {
        t.Fatalf("Failed to create config file: %v", err)
    }
    defer os.Remove(configFile)

    // Load the configuration
    cfg, err := LoadConfig()
    if err != nil {
        t.Fatalf("Failed to load config: %v", err)
    }

    // Verify the configuration
    expectedConfig := &Config{
        API: struct {
            URL   string `json:"url"`
            Token string `json:"token"`
        }{
            URL:   "https://api.example.com",
            Token: "testtoken",
        },
        Storage: struct {
            Path string `json:"path"`
        }{
            Path: "/tmp/backups",
        },
    }

    if !configEqual(cfg, expectedConfig) {
        t.Fatalf("Config mismatch: got %v, want %v", cfg, expectedConfig)
    }
}

func configEqual(a, b *Config) bool {
    return a.API.URL == b.API.URL &&
        a.API.Token == b.API.Token &&
        a.Storage.Path == b.Storage.Path
}
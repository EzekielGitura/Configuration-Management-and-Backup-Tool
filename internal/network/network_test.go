package network

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "networking-tool/pkg/api"
)

func TestGetNetworkInfo(t *testing.T) {
    // Create a test server
    testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/network/info" {
            t.Errorf("Expected path /network/info, got %s", r.URL.Path)
            http.Error(w, "Not Found", http.StatusNotFound)
            return
        }

        networkInfo := NetworkInfo{
            Status:  "active",
            Devices: []string{"eth0", "eth1"},
        }

        networkInfoJSON, err := json.Marshal(networkInfo)
        if err != nil {
            t.Errorf("Failed to marshal network info: %v", err)
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.Write(networkInfoJSON)
    }))
    defer testServer.Close()

    // Initialize API client
    apiClient := api.NewClient(testServer.URL, "testtoken")

    // Get network info
    networkInfo, err := GetNetworkInfo(apiClient)
    if err != nil {
        t.Fatalf("Failed to get network info: %v", err)
    }

    expectedNetworkInfo := `{"status":"active","devices":["eth0","eth1"]}`
    if networkInfo != expectedNetworkInfo {
        t.Fatalf("Network info mismatch: got %s, want %s", networkInfo, expectedNetworkInfo)
    }
}
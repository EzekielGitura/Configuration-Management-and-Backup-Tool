package network

import (
    "encoding/json"
    "fmt"
    "net/http"

    "networking-tool/pkg/api"
)

type NetworkInfo struct {
    Status  string   `json:"status"`
    Devices []string `json:"devices"`
}

func GetNetworkInfo(apiClient *api.Client) (string, error) {
    // Make a GET request to the API endpoint
    resp, err := apiClient.Get("/network/info")
    if err != nil {
        return "", fmt.Errorf("failed to get network info: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

    // Decode the response body into a NetworkInfo struct
    var networkInfo NetworkInfo
    err = json.NewDecoder(resp.Body).Decode(&networkInfo)
    if err != nil {
        return "", fmt.Errorf("failed to decode network info: %v", err)
    }

    // Convert NetworkInfo to JSON string
    networkInfoJSON, err := json.Marshal(networkInfo)
    if err != nil {
        return "", fmt.Errorf("failed to marshal network info: %v", err)
    }

    return string(networkInfoJSON), nil
}
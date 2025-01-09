package api

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestClient_Get(t *testing.T) {
    // Create a test server
    testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/test" {
            t.Errorf("Expected path /test, got %s", r.URL.Path)
            http.Error(w, "Not Found", http.StatusNotFound)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(`{"key": "value"}`))
    }))
    defer testServer.Close()

    // Initialize API client
    apiClient := NewClient(testServer.URL, "testtoken")

    // Make a GET request
    resp, err := apiClient.Get("/test")
    if err != nil {
        t.Fatalf("Failed to make GET request: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        t.Fatalf("Unexpected status code: %d", resp.StatusCode)
    }

    var result map[string]string
    err = json.NewDecoder(resp.Body).Decode(&result)
    if err != nil {
        t.Fatalf("Failed to decode response: %v", err)
    }

    expectedResult := map[string]string{"key": "value"}
    if result["key"] != expectedResult["key"] {
        t.Fatalf("Response mismatch: got %v, want %v", result, expectedResult)
    }
}

func TestClient_Post(t *testing.T) {
    // Create a test server
    testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/test" {
            t.Errorf("Expected path /test, got %s", r.URL.Path)
            http.Error(w, "Not Found", http.StatusNotFound)
            return
        }

        if r.Method != "POST" {
            t.Errorf("Expected POST request, got %s", r.Method)
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }

        var data map[string]string
        err := json.NewDecoder(r.Body).Decode(&data)
        if err != nil {
            t.Errorf("Failed to decode request body: %v", err)
            http.Error(w, "Bad Request", http.StatusBadRequest)
            return
        }

        expectedData := map[string]string{"key": "value"}
        if data["key"] != expectedData["key"] {
            t.Errorf("Request

			func TestClient_Post(t *testing.T) {
    // Create a test server
    testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/test" {
            t.Errorf("Expected path /test, got %s", r.URL.Path)
            http.Error(w, "Not Found", http.StatusNotFound)
            return
        }

        if r.Method != "POST" {
            t.Errorf("Expected POST request, got %s", r.Method)
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }

        var data map[string]string
        err := json.NewDecoder(r.Body).Decode(&data)
        if err != nil {
            t.Errorf("Failed to decode request body: %v", err)
            http.Error(w, "Bad Request", http.StatusBadRequest)
            return
        }

        expectedData := map[string]string{"key": "value"}
        if data["key"] != expectedData["key"] {
            t.Errorf("Request body mismatch: got %v, want %v", data, expectedData)
            http.Error(w, "Bad Request", http.StatusBadRequest)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(`{"status": "success"}`))
    }))
    defer testServer.Close()

    // Initialize API client
    apiClient := NewClient(testServer.URL, "testtoken")

    // Data to send in POST request
    data := map[string]string{"key": "value"}

    // Make a POST request
    resp, err := apiClient.Post("/test", data)
    if err != nil {
        t.Fatalf("Failed to make POST request: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        t.Fatalf("Unexpected status code: %d", resp.StatusCode)
    }

    var result map[string]string
    err = json.NewDecoder(resp.Body).Decode(&result)
    if err != nil {
        t.Fatalf("Failed to decode response: %v", err)
    }

    expectedResult := map[string]string{"status": "success"}
    if result["status"] != expectedResult["status"] {
        t.Fatalf("Response mismatch: got %v, want %v", result, expectedResult)
    }
}
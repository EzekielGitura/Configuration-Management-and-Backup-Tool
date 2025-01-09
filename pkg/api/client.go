package api

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

type Client struct {
    BaseURL string
    Token   string
}

func NewClient(baseURL, token string) *Client {
    return &Client{
        BaseURL: baseURL,
        Token:   token,
    }
}

func (c *Client) Get(endpoint string) (*http.Response, error) {
    url := fmt.Sprintf("%s%s", c.BaseURL, endpoint)
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to create request: %v", err)
    }

    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("failed to do request: %v", err)
    }

    return resp, nil
}

func (c *Client) Post(endpoint string, data interface{}) (*http.Response, error) {
    url := fmt.Sprintf("%s%s", c.BaseURL, endpoint)
    jsonData, err := json.Marshal(data)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal data: %v", err)
    }

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
    if err != nil {
        return nil, fmt.Errorf("failed to create request: %v", err)
    }

    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("failed to do request: %v", err)
    }

    return resp, nil
}
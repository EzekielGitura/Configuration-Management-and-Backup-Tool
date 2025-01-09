package config

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "os"
)

type Config struct {
    API struct {
        URL   string `json:"url"`
        Token string `json:"token"`
    } `json:"api"`
    Storage struct {
        Path string `json:"path"`
    } `json:"storage"`
}

func LoadConfig() (*Config, error) {
    // Read the configuration file
    configFile := "config.json"
    data, err := ioutil.ReadFile(configFile)
    if err != nil {
        return nil, fmt.Errorf("failed to read config file: %v", err)
    }

    // Unmarshal the JSON data into a Config struct
    var cfg Config
    err = json.Unmarshal(data, &cfg)
    if err != nil {
        return nil, fmt.Errorf("failed to unmarshal config: %v", err)
    }

    return &cfg, nil
}
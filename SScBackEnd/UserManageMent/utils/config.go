package utils

import (
    "encoding/json"
    "os"
)

// Config struct to hold the configuration settings
type Config struct {
    Server struct {
        Port string `json:"port"`
        IP   string `json:"ip"`
    } `json:"server"`
}

// LoadConfig function reads configuration from a JSON file
func LoadConfig(path string) (*Config, error) {
    var config Config

    // Open the configuration file
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    // Decode JSON file into Config struct
    decoder := json.NewDecoder(file)
    if err := decoder.Decode(&config); err != nil {
        return nil, err
    }

    return &config, nil
}

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Config holds the configuration values for the client.
type Config struct {
	Database struct {
		IP   string `json:"ip"`
		Port int    `json:"port"`
	} `json:"database"`
}

// Client struct holds the configuration and other relevant data.
type Client struct {
	Config Config
}

// NewClient initializes a new Client with the configuration read from a file.
func NewClient() (*Client, error) {
	configFile:= "../config/userConfig.json"
	config, err := readConfig(configFile)
	if err != nil {
		return nil, err
	}

	return &Client{
		Config: config,
	}, nil
}

// readConfig reads and parses the configuration file.
func readConfig(filename string) (Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Config{}, fmt.Errorf("error opening config file: %v", err)
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, fmt.Errorf("error decoding config file: %v", err)
	}

	// Convert port from string to int
	port:= config.Database.Port
	if err != nil {
		return Config{}, fmt.Errorf("error converting port to int: %v", err)
	}
	config.Database.Port = port

	return config, nil
}

// buildURL constructs the full URL for a given path.
func (c *Client) buildURL(path string) string {
	return fmt.Sprintf("http://%s:%d%s", c.Config.Database.IP, c.Config.Database.Port, path)
}

// getUserByUserName sends a request to get a user by username.
// Path: /v1/database/user/{username}
func (c *Client) getUserByUserName(username string) (*http.Response, error) {
	path := fmt.Sprintf("/v1/database/user/%s", username)
	fmt.Printf("Request URL: %s\n", c.buildURL(path))
	resp, err := http.Get(c.buildURL(path))
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	return resp, nil
}

// createUser sends a request to create a new user.
// Path: /v1/database/user
func (c *Client) createUser(userData []byte) (*http.Response, error) {
	path := "/v1/database/user"
	fmt.Printf("Request URL: %s\n", c.buildURL(path))
	resp, err := http.Post(c.buildURL(path), "application/json", ioutil.NopCloser(bytes.NewReader(userData)))
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	return resp, nil
}

// updateUser sends a request to update an existing user.
// Path: /v1/database/user/{username}
func (c *Client) updateUser(username string, userData []byte) (*http.Response, error) {
	path := fmt.Sprintf("/v1/database/user/%s", username)
	fmt.Printf("Request URL: %s\n", c.buildURL(path))
	req, err := http.NewRequest(http.MethodPut, c.buildURL(path), ioutil.NopCloser(bytes.NewReader(userData)))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	return resp, nil
}

// deleteUser sends a request to delete a user by username.
// Path: /v1/database/user/{username}
func (c *Client) deleteUser(username string) (*http.Response, error) {
	path := fmt.Sprintf("/v1/database/user/%s", username)
	fmt.Printf("Request URL: %s\n", c.buildURL(path))
	req, err := http.NewRequest(http.MethodDelete, c.buildURL(path), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	return resp, nil
}

// listUsers sends a request to list all users.
// Path: /v1/database/users
func (c *Client) listUsers() (*http.Response, error) {
	path := "/v1/database/users"
	fmt.Printf("Request URL: %s\n", c.buildURL(path))
	resp, err := http.Get(c.buildURL(path))
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	return resp, nil
}

package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchLocations(url string, config *Config) error {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error making GET request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("bad status: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	err = json.Unmarshal(body, config)
	if err != nil {
		return fmt.Errorf("error unmarshaling JSON: %w", err)
	}
	return nil
}

func HandleMapCommand(config *Config) error {
	err := FetchLocations(config.Next, config)
	if err != nil {
		return fmt.Errorf("error fetching locations: %w", err)
	}

	for _, result := range config.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func HandleMapBackCommand(config *Config) error {
	if config.Previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	err := FetchLocations(config.Previous, config)
	if err != nil {
		return fmt.Errorf("error fetching locations: %w", err)
	}

	for _, result := range config.Results {
		fmt.Println(result.Name)
	}
	return nil
}

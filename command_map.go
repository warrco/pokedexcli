package main

import (
	"encoding/json"
	"fmt"

	"github.com/warrco/pokedexcli/internal/pokeapi"
)

func HandleMapCommand(config *pokeapi.Locations) error {

	if config.Next == "" {
		fmt.Println("You're on the last page")
		return nil
	}

	currentURL := config.Next

	if data, ok := cache.Get(currentURL); ok {
		fmt.Println("Cache hit!")
		err := json.Unmarshal(data, config)
		if err != nil {
			return fmt.Errorf("error unmarshaling JSON: %w", err)
		}
	} else {
		fmt.Println("Cache miss, fetching from API...")
		err := pokeapi.FetchLocations(currentURL, config)
		if err != nil {
			return fmt.Errorf("error fetching locations: %w", err)
		}

		data, err := json.Marshal(config)
		if err != nil {
			return fmt.Errorf("error marshaling JSON: %w", err)
		}
		cache.Add(currentURL, data)
	}

	for _, result := range config.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func HandleMapBackCommand(config *pokeapi.Locations) error {
	if config.Previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	currentURL := config.Previous

	if data, ok := cache.Get(currentURL); ok {
		fmt.Println("Cache hit!")
		err := json.Unmarshal(data, config)
		if err != nil {
			return fmt.Errorf("error unmarshaling JSON: %w", err)
		}
	} else {
		fmt.Println("Cache miss, fetching from API...")

		err := pokeapi.FetchLocations(currentURL, config)
		if err != nil {
			return fmt.Errorf("error fetching locations: %w", err)
		}

		data, err := json.Marshal(config)
		if err != nil {
			return fmt.Errorf("error marshaling JSON: %w", err)
		}
		cache.Add(currentURL, data)
	}

	for _, result := range config.Results {
		fmt.Println(result.Name)
	}
	return nil
}

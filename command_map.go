package main

import (
	"fmt"
)

func commandMapf(cfg *config) error {
	areaLocations, err := cfg.pokeapiClient.FetchLocations(cfg.nextLocationsURL)
	if err != nil {
		return fmt.Errorf("failed to fetch locations: %w", err)
	}

	cfg.nextLocationsURL = areaLocations.Next
	cfg.prevLocationsURL = areaLocations.Previous

	for _, result := range areaLocations.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		fmt.Println("you are on the first page")
		return nil
	}

	areaLocations, err := cfg.pokeapiClient.FetchLocations(cfg.prevLocationsURL)
	if err != nil {
		return fmt.Errorf("error fetching locations: %w", err)
	}

	cfg.nextLocationsURL = areaLocations.Next
	cfg.prevLocationsURL = areaLocations.Previous

	for _, result := range areaLocations.Results {
		fmt.Println(result.Name)
	}
	return nil
}

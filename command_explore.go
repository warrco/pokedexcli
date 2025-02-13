package main

import (
	"fmt"
)

func commandExplore(cfg *config) error {
	if len(cfg.parameters) == 0 {
		return fmt.Errorf("please provide a valid location")
	}
	locationName := cfg.parameters[0]

	pokemonInfo, err := cfg.pokeapiClient.FetchLocationInfo(locationName)
	if err != nil {
		return fmt.Errorf("could not retrieve location: %w", err)
	}

	fmt.Printf("Exploring %s...\n", locationName)
	fmt.Println("Found Pokemon:")

	for _, p := range pokemonInfo.PokemonEncounters {
		fmt.Printf(" - %s\n", p.Pokemon.Name)
	}
	return nil
}

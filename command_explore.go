package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("please provide a valid location")
	}
	locationName := args[0]

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

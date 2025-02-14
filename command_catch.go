package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you threw your pokeball into the void")
	}
	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.FetchPokemon(pokemonName)
	if err != nil {
		return fmt.Errorf("error retrieving pokemon name: %w", err)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	threshold := 70 - (pokemon.BaseExperience / 10)
	chanceToCatch := rand.Intn(100)
	if chanceToCatch < threshold {
		cfg.caughtPokemon[pokemon.Name] = pokemon
		fmt.Printf("%s was caught and added to your Pokedex!\n", pokemonName)
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	return nil
}

package main

import (
	"fmt"
	"math/rand"

	"github.com/warrco/pokedexcli/internal/pokeapi"
)

func addToPokedex(pokemon pokeapi.Pokemon) {
	pokedex[pokemon.Name] = pokemon
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you threw your pokeball into the void")
	}
	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.FetchPokemon(pokemonName)
	if err != nil {
		return fmt.Errorf("error retrieving pokemon name: %w", err)
	}

	fmt.Printf("Throwing pokeball at %s\n", pokemonName)

	threshold := 70 - (pokemon.BaseExperience / 10)
	chanceToCatch := rand.Intn(100)
	if chanceToCatch < threshold {
		addToPokedex(pokemon)
		fmt.Printf("%s was caught and added to your Pokedex!\n", pokemonName)
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	return nil
}

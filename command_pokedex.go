package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for key := range cfg.caughtPokemon {
		fmt.Printf(" - %v\n", key)
	}
	return nil
}

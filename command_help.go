package main

import (
	"fmt"

	"github.com/warrco/pokedexcli/internal/pokeapi"
)

func commandHelp(_ *pokeapi.Locations) error {

	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for key, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", key, cmd.description)
	}
	fmt.Println()

	return nil
}

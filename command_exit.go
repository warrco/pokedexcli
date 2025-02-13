package main

import (
	"fmt"
	"os"

	"github.com/warrco/pokedexcli/internal/pokeapi"
)

func commandExit(_ *pokeapi.Locations) error {

	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

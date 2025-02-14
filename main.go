package main

import (
	"time"

	"github.com/warrco/pokedexcli/internal/pokeapi"
)

var pokedex = make(map[string]pokeapi.Pokemon)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/warrco/pokedexcli/internal/pokeapi"
	"github.com/warrco/pokedexcli/internal/pokecache"
)

var cache *pokecache.Cache

func startRepl() {
	cache = pokecache.NewCache(5 * time.Second)
	scanner := bufio.NewScanner(os.Stdin)
	var config pokeapi.Locations

	err := pokeapi.FetchLocations("", &config)
	if err != nil {
		fmt.Println("error initializing locations", err)
		return
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		command := cleanInput(input)
		if len(command) == 0 {
			continue
		}

		commandName := command[0]
		if cmd, exists := getCommands()[commandName]; exists {
			err := cmd.callback(&config)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}

func cleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	words := strings.Fields(lowerCase)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Locations) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays map locations",
			callback:    HandleMapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous map locations",
			callback:    HandleMapBackCommand,
		},
	}
}

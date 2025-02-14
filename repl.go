package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/warrco/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	parameters       []string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		command := cleanInput(input)
		if len(command) == 0 {
			continue
		}

		commandName := command[0]
		args := []string{}
		if len(command) > 1 {
			args = command[1:]
		}
		cfg.parameters = args

		if cmd, exists := getCommands()[commandName]; exists {
			err := cmd.callback(cfg, args...)
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
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{

		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays map locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous map locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Search an area for pokemon",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/warrco/pokedexcli/internal/pokeapi"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	var config pokeapi.Config
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
	callback    func(*pokeapi.Config) error
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
			callback:    pokeapi.HandleMapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous map locations",
			callback:    pokeapi.HandleMapBackCommand,
		},
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var cfg *Config = &Config{}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		commandArgs := words[1:]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(commandArgs...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Previous location areas",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "explore pokemon location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "catch the pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspect the pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "list all pokemons",
			callback:    commandPokedex,
		},
	}
}

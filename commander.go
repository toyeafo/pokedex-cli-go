package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/toyeafo/pokedex-cli-go/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationURL     *string
	previousLocationURL *string
}

func startComm(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()

		words := cleanUserInput(scanner.Text())
		if len(words) == 0 {
			fmt.Println("Please Enter a command")
			continue
		}

		commandName := words[0]
		var areaName string
		if len(words) > 1 {
			areaName = words[1]
		}

		command, exists := commandInp()[commandName]
		if exists {
			err := command.callback(cfg, areaName)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanUserInput(text string) []string {
	output := strings.ToLower((text))
	words := strings.Fields(output)
	return words
}

func commandInp() map[string]cliCommand {
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
			description: "Displays names of 20 location areas in Pokemon World",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays more information about a location area",
			callback:    commandExplore,
		},
	}
}

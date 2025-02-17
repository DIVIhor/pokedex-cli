package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/DIVIgor/pokedex-cli/internal/pokeAPI"
)

// cli command template
type cliCommand struct {
	name string
	description string
	callback func(*apiConfig) error
}

type apiConfig struct {
    client pokeAPI.Client
    next string
    previous string
}


// cli commands registry
func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: exitRepl,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: usageHelp,
		},
		"map": {
			name: "map",
			description: "Displays the next 20 locations on the map",
			callback: mapNext,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the previous 20 locations on the map",
			callback: mapPrev,
		},
	}
}


// Clean an inputted string from spaces and split it to a slice of strings
func cleanInput(text string) (cleanedInput []string) {
	if len(text) > 0 {
		input := strings.ToLower(text)
		cleanedInput = strings.Fields(input)
	}

	return cleanedInput
}

func startRepl(cfg *apiConfig) {
	cliCommands := getCommands()
	// wait for user input
	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if !input.Scan() {
			continue
		}

		inputStr := input.Text()
		words := cleanInput(inputStr)

		command, exists := cliCommands[words[0]]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}
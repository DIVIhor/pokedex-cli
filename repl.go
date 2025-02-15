package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// cli command template
type cliCommand struct {
	name string
	description string
	callback func() error
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

func startRepl() {
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
		err := command.callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}
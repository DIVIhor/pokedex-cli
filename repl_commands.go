package main

import (
	"fmt"
	"os"
)

// print the exit message and close the app
func exitRepl() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

// show command usage info
func usageHelp() error {
	var info string
	prefix := "Welcome to the Pokedex!\nUsage:\n\n%s"
	infoTemplate := "%s: %s\n"

	cliCommands := getCommands()
	for _, command := range cliCommands {
		info += fmt.Sprintf(infoTemplate, command.name, command.description)
	}

	fmt.Printf(prefix, info)
	return nil
}
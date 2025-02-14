package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Clean an inputted string from spaces and split it to a slice of strings
func cleanInput(text string) (cleanedInput []string) {
	if len(text) > 0 {
		input := strings.ToLower(text)
		cleanedInput = strings.Fields(input)
	}

	return cleanedInput
}

func startRepl() {
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !input.Scan() {
			continue
		}
		str := input.Text()
		strs := cleanInput(str)
		fmt.Printf("Your command was: %s\n", strs[0])
	}
}
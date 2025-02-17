package main

import (
	"fmt"
	"os"

	"github.com/DIVIgor/pokedex-cli/internal/pokeAPI"
)

// print the exit message and close the app
func exitRepl(api *apiConfig) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

// show command usage info
func usageHelp(api *apiConfig) error {
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

// get and print locations for new location chunk
func locationProcessor(url string, api *apiConfig) (err error) {
	location, err := pokeAPI.GetLocations(url)
    if err != nil {return}

    api.next = location.Next
    api.previous = location.Previous

    for _, loc := range location.Results {
        fmt.Println(loc.Name)
    }
	return
}

// get map for the next location with caching
func mapNext(api *apiConfig) (err error) {
	url := api.next
    if len(url) == 0 {
        url = "https://pokeapi.co/api/v2/location-area/"
    }
    
	err = locationProcessor(url, api)

	return
}

// get map for the previous location with caching
func mapPrev(api *apiConfig) (err error) {
	url := api.previous
	if len(url) == 0 {
		fmt.Println("you're on the first page")
		return
	}

	err = locationProcessor(url, api)
	
	return
}
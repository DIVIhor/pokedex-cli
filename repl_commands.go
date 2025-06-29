package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/DIVIhor/pokedex-cli/internal/pokeAPI"
)

const areaURL string = "https://pokeapi.co/api/v2/location-area/"

// print the exit message and close the app
func exitRepl(api *apiConfig, locName string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

// show command usage info
func usageHelp(api *apiConfig, locName string) error {
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

// get and print locations for a new/cached map chunk
func locationProcessor(url string, api *apiConfig) (err error) {
	data, err := api.client.GetRawData(url)
	if err != nil {
		return
	}

	location, err := pokeAPI.ReadLocationResp(data)
	if err != nil {
		return err
	}

	api.next = location.Next
	api.previous = location.Previous

	for _, loc := range location.Results {
		fmt.Println(loc.Name)
	}

	return
}

// get map for the next location with caching
func mapNext(api *apiConfig, locName string) (err error) {
	url := api.next
	if len(url) == 0 {
		url = areaURL //"https://pokeapi.co/api/v2/location-area/"
	}

	err = locationProcessor(url, api)

	return
}

// get map for the previous location with caching
func mapPrev(api *apiConfig, locName string) (err error) {
	url := api.previous
	if len(url) == 0 {
		fmt.Println("you're on the first page")
		return
	}

	err = locationProcessor(url, api)

	return
}

func locDetailsProcessor(url string, api *apiConfig) (err error) {
	data, err := api.client.GetRawData(url)
	if err != nil {
		return
	}

	details, err := pokeAPI.ReadDetailsResp(data)
	if err != nil {
		return err
	}

	// since the location could be inputted as an id
	// the location name should be printed from the parsed response
	fmt.Printf("Exploring %s...\n", details.Name)

	fmt.Println("Found Pokemon:")
	for _, encounter := range details.PokemonEncounters {
		fmt.Printf("  - %s\n", encounter.Pokemon.Name)
	}

	return
}

// explore the location for pokemon presence
func explore(api *apiConfig, locName string) (err error) {
	if len(locName) == 0 {
		return errors.New("the location cannot be empty")
	}

	url := areaURL + locName
	err = locDetailsProcessor(url, api)

	return
}

func pokemonProcessor(pokemonName string, api *apiConfig) (err error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokemonName)
	data, err := api.client.GetRawData(url)
	if err != nil {
		return
	}

	pokemon, err := pokeAPI.ReadPokemonResp(data)
	if err != nil {
		return
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	// determine a chance on catching the pokemon
	minCatchChance := 0.8
	chance := float64(rand.Intn(pokemon.BaseExperience)) / float64(pokemon.BaseExperience)
	if chance < minCatchChance {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return
	}
	// in case of catching -> add it to the user's pokedex
	fmt.Printf("%s was caught!\n", pokemon.Name)
	api.caughtPokemon[pokemonName] = pokemon
	fmt.Println("You may now inspect it with the 'inspect' command.")

	return
}

// try to catch a pokemon
func catch(api *apiConfig, pokemonName string) (err error) {
	if len(pokemonName) == 0 {
		return errors.New("the Pokemon name cannot be empty")
	}

	err = pokemonProcessor(pokemonName, api)
	return
}

// inspect a caught pokemon
func inspect(api *apiConfig, pokemonName string) (err error) {
	if len(pokemonName) == 0 {
		return errors.New("the Pokemon name cannot be empty")
	}

	pokemon, exists := api.caughtPokemon[strings.ToLower(pokemonName)]
	if !exists {
		return fmt.Errorf("you haven't caught %s", pokemonName)
	}

	fmt.Printf(
		`Name: %s
Height: %d
Weight: %d
Stats:`, pokemon.Name, pokemon.Height, pokemon.Weight)
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pType := range pokemon.Types {
		fmt.Println("  -", pType.Type.Name)
	}

	return
}

// show all the caught pokemon
func showPokedex(api *apiConfig, _ string) (err error) {
	if len(api.caughtPokemon) == 0 {
		return errors.New("you have not caught any Pokemon")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range api.caughtPokemon {
		fmt.Println("  -", pokemon.Name)
	}

	return
}

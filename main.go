package main

import (
	"time"

	"github.com/DIVIhor/pokedex-cli/internal/pokeAPI"
)

func main() {
	pokeClient := pokeAPI.NewClient(5*time.Second, 5*time.Minute)
	apiCfg := &apiConfig{
		client:        pokeClient,
		caughtPokemon: map[string]pokeAPI.PokemonResponse{},
	}

	startRepl(apiCfg)
}

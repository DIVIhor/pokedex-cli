package main

import (
	"time"

	"github.com/DIVIgor/pokedex-cli/internal/pokeAPI"
)


func main() {
	pokeClient := pokeAPI.NewClient(5 * time.Second)
	apiCfg := &apiConfig{
		client: pokeClient,
	}

	startRepl(apiCfg)
}

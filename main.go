package main

import (
	"time"

	"github.com/binaryCoder-101/pokedexcli/internal/pokeapi"
	// "github.com/binaryCoder-101/pokedexcli/internal/pokecache"
)

func main() {
	pokeAPIClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		httpClient: pokeAPIClient,
		prev:       nil,
		next:       nil,
	}
	// pokecache.NewCache(5 * time.Second)
	startRepl(cfg)
}

package main

import (
	"time"

	"github.com/binaryCoder-101/pokedexcli/internal/pokeapi"
	"github.com/binaryCoder-101/pokedexcli/internal/pokecache"
)

func main() {
	pokeAPIClient := pokeapi.NewClient(5 * time.Second)
	cache := pokecache.NewCache(30 * time.Second)

	cfg := &config{
		httpClient: pokeAPIClient,
		cache:      cache,
		prev:       nil,
		next:       nil,
	}
	startRepl(cfg)
}

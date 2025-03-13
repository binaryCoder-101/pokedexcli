package main

import (
	"time"

	"github.com/binaryCoder-101/pokedexcli/internal/pokeapi"
	"github.com/binaryCoder-101/pokedexcli/internal/pokecache"
	"github.com/binaryCoder-101/pokedexcli/internal/pokedex"
)

func main() {
	pokeAPIClient := pokeapi.NewClient(5 * time.Second)
	cache := pokecache.NewCache(30 * time.Second)
	pokedexDatabase := pokedex.NewPokedex()

	cfg := &config{
		httpClient:      pokeAPIClient,
		pokedexDatabase: pokedexDatabase,
		cache:           cache,
		prev:            nil,
		next:            nil,
	}
	startRepl(cfg)
}

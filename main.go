package main

import (
	"time"

	"github.com/binaryCoder-101/pokedexcli/internal/pokecache"
)

func main() {
	cfg := &config{}
	pokecache.NewCache(5 * time.Second)
	startRepl(cfg)
}

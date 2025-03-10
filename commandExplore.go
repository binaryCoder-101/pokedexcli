package main

import (
	"fmt"

	"github.com/binaryCoder-101/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) > 1 {
		return fmt.Errorf("too many arguments")
	}
	if len(args) < 1 {
		return fmt.Errorf("too few arguments")
	}

	urlInput := pokeapi.BaseURL + args[0]
	var sliceofBytes []byte

	value, exists := cfg.cache.Get(urlInput)
	if exists {
		sliceofBytes = value
	} else {
		respData, err := cfg.httpClient.ResponseData(&urlInput)
		if err != nil {
			return err
		}
		sliceofBytes = respData
		cfg.cache.Add(urlInput, sliceofBytes)
	}

	locationAreasDetails, err := pokeapi.UnmarshalSliceOfBytesLocationAreasDetails(sliceofBytes)
	if err != nil {
		return err
	}

	for _, locationAreasDetail := range locationAreasDetails.PokemonEncounters {
		fmt.Println(locationAreasDetail.Pokemon.Name)
	}

	return nil
}

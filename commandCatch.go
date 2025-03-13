package main

import (
	"fmt"
	"math/rand"

	"github.com/binaryCoder-101/pokedexcli/internal/pokeapi"
	"github.com/binaryCoder-101/pokedexcli/internal/pokedex"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) > 1 {
		return fmt.Errorf("too many arguments")
	}
	if len(args) < 1 {
		return fmt.Errorf("too few arguments")
	}

	pokemonName := args[0]

	if _, exists := cfg.pokedexDatabase.CaughtPokemon[pokemonName]; exists {

		fmt.Println("you already caught this pokémon")
		return nil

	} else {

		urlInput := pokeapi.BaseURL + "/pokemon/" + pokemonName

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

		pokemonDetails, err := pokeapi.UnmarshalSliceOfBytesPokemonDetails(sliceofBytes)
		if err != nil {
			return err
		}

		baseExperience := pokemonDetails.BaseExperience

		fmt.Println("Throwing a Pokeball at " + pokemonName + "...")

		// cfg.pokedexDatabase = &pokedex.PokedexStruct{
		// 	CaughtPokemon: map[string]pokedex.Pokemon{
		// 		pokemonName: {
		// 			Name:           pokemonName,
		// 			BaseExperience: baseExperience,
		// 		},
		// 	},
		// }

		threshold := baseExperience / 2
		prob := rand.Intn(baseExperience)

		if prob > threshold {
			fmt.Println("Pokemon escaped!")
			return nil
		} else {
			cfg.pokedexDatabase.CaughtPokemon[pokemonName] = pokedex.Pokemon{
				Name:           pokemonName,
				BaseExperience: baseExperience,
			}

			fmt.Println(pokemonName + " was caught!")
			return nil
		}

	}
}
